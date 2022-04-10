create table tiktok.Video
(
    id          char(20)     null,
    description text         null,
    views       int          null,
    comments    int          null,
    shares      int          null,
    likes       int          null,
    postedBy    varchar(128) null,
    private     bool         null,
    added       date         null,
    constraint Video_pk
        primary key (id),
    constraint Video_Account_username_fk
        foreign key (postedBy) references tiktok.Account (username)
);

