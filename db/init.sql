-- Table creation

create table if not exists public.users(
    id bigint unique
);

create table if not exists public.segments(
    id bigint generated always as identity primary key,
    name text unique not null
);

create table if not exists public.users_and_segments(
    user_id bigint,
    foreign key (user_id) references public.users(id),
    segment_id bigint,
    foreign key (segment_id) references public.segments(id),
    expiration_date timestamp default null,
    constraint unique_linkage unique (user_id, segment_id)
);

create table if not exists public.user_history(
    user_id bigint,
    segment_name text not null,
    operation text not null,
    registration_date timestamp default now()
);

-- Managing links between users and segments

create or replace procedure public.insert_user_segments(names text[], uid bigint, expire timestamp)
as $$
declare segments_id bigint[];
begin
    select array_agg(s.id)
    from public.segments s
    where s.name = any(names)
    into segments_id;

    insert into public.users_and_segments(user_id, segment_id, expiration_date)
    values (uid, unnest(segments_id), expire)
    on conflict (user_id, segment_id) do update set expiration_date = expire;
end
$$ language plpgsql;

create or replace procedure public.delete_user_segments(names text[], uid bigint)
as $$
declare segments_id bigint[];
begin
    select array_agg(s.id)
    from public.segments s
    where s.name = any(names)
    into segments_id;

    delete from public.users_and_segments us
    where us.user_id = uid and us.segment_id = any(segments_id);
end
$$ language plpgsql;

create or replace procedure public.clear_expired_linkages()
as $$
begin
    delete from public.users_and_segments
    where expiration_date is not null and expiration_date < now();
end
$$ language plpgsql;

create or replace function public.select_user_segments(uid bigint)
returns table (
    name text
)
as $$
begin
    call public.clear_expired_linkages();

    return query
    select s.name
    from public.segments s join public.users_and_segments us
        on s.id = us.segment_id and us.user_id = uid;
end
$$ language plpgsql;

-- Managing segments

create or replace function public.select_rand_users_id(user_percentage float)
returns table (
    user_id bigint
)
as $$
declare users_count bigint;
begin
    select count(u.id) * user_percentage / 100.0
    from public.users u
    into users_count;

    return query
    select u.id
    from public.users u
    limit users_count;
end
$$ language plpgsql;

create or replace procedure public.insert_segment(seg_name text, user_percentage float, expire timestamp)
as $$
declare seg_id bigint;
begin
    select -1 into seg_id;

    insert into public.segments(name)
    values (seg_name)
    on conflict (name) do nothing
    returning id into seg_id;

    if seg_id != -1 and abs(user_percentage) > 0.0001 then
        insert into public.users_and_segments(user_id, segment_id, expiration_date)
        values (select_rand_users_id(user_percentage), seg_id, expire)
        on conflict (user_id, segment_id) do update set expiration_date = expire;
    end if;
end
$$ language plpgsql;

create or replace procedure public.delete_segment(seg_name text)
as $$
begin
    delete from public.segments s
    where s.name = seg_name;
end
$$ language plpgsql;

create or replace function delete_users_segment_trigger()
returns trigger
as $$
begin
    raise notice 'Old =  %', old;
    raise notice 'New =  %', new;

    delete from public.users_and_segments us
    where us.segment_id = old.id;
    return old;
end;
$$ language plpgsql;

create or replace trigger delete_segment
before delete on public.segments
for each row execute procedure delete_users_segment_trigger();

create or replace function delete_user_segments_trigger()
returns trigger
as $$
begin
    raise notice 'Old =  %', old;
    raise notice 'New =  %', new;

    delete from public.users_and_segments us
    where us.user_id = new.id;
    return new;
end;
$$ language plpgsql;

create or replace trigger delete_user
after delete on public.users
for each row execute procedure delete_user_segments_trigger();

-- Managing history

create or replace function users_and_segments_insert_trigger()
returns trigger
as $$
declare segment_names table (
    name text
);
begin
    raise notice 'Old =  %', old;
    raise notice 'New =  %', new;

    select s.name
    from public.segments s
    where s.id = new.segment_id
    into segment_names;

    insert into public.user_history(user_id, segment_name, operation)
    values (new.user_id, segment_names.name, 'insert');
    return new;
end;
$$ language plpgsql;

create or replace function users_and_segments_delete_trigger()
    returns trigger
as $$
declare segment_names table (
    name text
);
begin
    raise notice 'Old =  %', old;
    raise notice 'New =  %', new;

    select s.name
    from public.segments s
    where s.id = new.segment_id
    into segment_names;

    insert into public.user_history(user_id, segment_name, operation)
    values (new.user_id, segment_names.name, 'insert');
    return new;
end;
$$ language plpgsql;

create or replace trigger insert_user_and_segment_linkage
after insert on public.users_and_segments
for each row execute procedure users_and_segments_insert_trigger();

create or replace trigger delete_user_and_segment_linkage
after delete on public.users_and_segments
for each row execute procedure users_and_segments_delete_trigger();
