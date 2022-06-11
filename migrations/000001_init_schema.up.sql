CREATE TABLE organizations
(
    id               BIGSERIAL PRIMARY KEY,
    created_at       TIMESTAMP WITH TIME zone NOT NULL,
    updated_at       TIMESTAMP WITH TIME zone,
    name             VARCHAR(255),
    full_name        VARCHAR(255),
    about            TEXT,
    logo             VARCHAR(100),
    color            VARCHAR(40)[]            NOT NULL,
    url              VARCHAR(200),
    employees_amount INT                      NOT NULL,
    revenue          BIGINT                   NOT NULL,
    organization_inn VARCHAR(12)              NOT NULL
        CONSTRAINT organizations_organization_inn_uniq UNIQUE,
    phone            VARCHAR(30),
    email            VARCHAR(254)             NOT NULL,
    status           VARCHAR(50)              NOT NULL,
    gr_points        INT                      NOT NULL
);

CREATE TABLE organization_sectors
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME zone NOT NULL,
    updated_at TIMESTAMP WITH TIME zone,
    name       VARCHAR(255)             NOT NULL
);


CREATE TABLE organization_sectors_rel
(
    organization_id BIGINT NOT NULL
        CONSTRAINT organizationdb_id_fk_organizat REFERENCES organizations deferrable initially deferred,
    sectid          BIGINT NOT NULL
        CONSTRAINT organizationsectordb_fk_organizat REFERENCES organization_sectors deferrable initially deferred,
    CONSTRAINT organi_organizationdb_id_organi_uniq UNIQUE (organization_id, sectid)
);


CREATE TABLE organization_question_templates
(
    id               BIGSERIAL PRIMARY KEY,
    created_at       TIMESTAMP WITH TIME zone NOT NULL,
    updated_at       TIMESTAMP WITH TIME zone,
    text             VARCHAR(255)             NOT NULL,
    answer_type      VARCHAR(50)              NOT NULL,
    is_answer        boolean                  NOT NULL,
    is_required      boolean                  NOT NULL,
    is_private       boolean                  NOT NULL,
    opened_answer_id BIGINT
        CONSTRAINT opened_answer_id_fk_organizat
            REFERENCES organization_question_templates deferrable initially deferred,
    parent_id        BIGINT
        CONSTRAINT parent_id_fk_organizat
            REFERENCES organization_question_templates deferrable initially deferred,
    gr_points        FLOAT                    NOT NULL
);



CREATE TABLE organization_forms
(
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMP WITH TIME zone NOT NULL,
    updated_at      TIMESTAMP WITH TIME zone,
    text_answer     VARCHAR(255),
    organization_id BIGINT
        CONSTRAINT organization_id_fk_organizat
            REFERENCES organizations deferrable initially deferred,
    question_id     BIGINT                   NOT NULL
        CONSTRAINT question_id_fk_organizat
            REFERENCES organization_question_templates deferrable initially deferred
);


CREATE TABLE organization_forms_answers
(
    organization_form_id     BIGINT NOT NULL
        CONSTRAINT organizationforms_i_fk_organizat
            REFERENCES organization_forms deferrable initially deferred,
    organization_template_id BIGINT NOT NULL
        CONSTRAINT organizationtemplate_fk_organizat
            REFERENCES organization_question_templates deferrable initially deferred,
    CONSTRAINT organizationforms_id_uniq UNIQUE (organization_form_id, organization_template_id)
);


