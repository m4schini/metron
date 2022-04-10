create database tiktok DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
;

create table tiktok.Account
(
    username    varchar(128) null,
    displayname varchar(128) null,
    bio         text         null,
    following   int          null,
    followers   int          null,
    likes       int          null,
    constraint Account_pk
        primary key (username)
) WITH SYSTEM VERSIONING;

create table tiktok.Video
(
    id          char(20)     null,
    description text         null,
    views       int          null,
    comments    int          null,
    shares      int          null,
    likes       int          null,
    videoLength int null,
    postedBy    varchar(128) null,
    available     bool         null,
    added       datetime     null,

    constraint Video_pk
        primary key (id),
    constraint Video_Account_username_fk
        foreign key (postedBy) references tiktok.Account (username)
) WITH SYSTEM VERSIONING;

create table tiktok.Tag
(
    name varchar(64) null,
    id   int auto_increment,
    constraint Tag_pk
        primary key (id),
    constraint Tag_name_uindex
        unique (name)
) WITH SYSTEM VERSIONING;

create table tiktok.video_mentions_Account
(
    post_ID          char(20)     not null,
    account_Username varchar(128) not null
);

create table tiktok.Video_has_Tag
(
    post_ID char(20) null,
    tag_ID  int      null
);

