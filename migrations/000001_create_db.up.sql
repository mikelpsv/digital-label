CREATE TABLE data_links
(
    key_link varchar(36)  default ''::character varying not null
        constraint data_links_pk
            primary key,
    key_data varchar(150) default ''::character varying,
    created_at timestamptz,
    data_type     integer      default 0                     not null,
    payload  jsonb                                      not null
);

ALTER TABLE data_links
    OWNER TO dl_user;

INSERT INTO data_links (key_link, created_at, payload) VALUES ('123', now(), '{}');