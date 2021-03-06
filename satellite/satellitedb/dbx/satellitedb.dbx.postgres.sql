-- AUTOGENERATED BY gopkg.in/spacemonkeygo/dbx.v1
-- DO NOT EDIT
CREATE TABLE accounting_raws (
	id bigserial NOT NULL,
	node_id bytea NOT NULL,
	interval_end_time timestamp with time zone NOT NULL,
	data_total double precision NOT NULL,
	data_type integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_rollups (
	id bigserial NOT NULL,
	node_id bytea NOT NULL,
	start_time timestamp with time zone NOT NULL,
	put_total bigint NOT NULL,
	get_total bigint NOT NULL,
	get_audit_total bigint NOT NULL,
	get_repair_total bigint NOT NULL,
	put_repair_total bigint NOT NULL,
	at_rest_total double precision NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_timestamps (
	name text NOT NULL,
	value timestamp with time zone NOT NULL,
	PRIMARY KEY ( name )
);
CREATE TABLE bucket_bandwidth_rollups (
	bucket_name bytea NOT NULL,
	project_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	interval_seconds integer NOT NULL,
	action integer NOT NULL,
	inline bigint NOT NULL,
	allocated bigint NOT NULL,
	settled bigint NOT NULL,
	PRIMARY KEY ( bucket_name, project_id, interval_start, action )
);
CREATE TABLE bucket_storage_tallies (
	bucket_name bytea NOT NULL,
	project_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	inline bigint NOT NULL,
	remote bigint NOT NULL,
	remote_segments_count integer NOT NULL,
	inline_segments_count integer NOT NULL,
	object_count integer NOT NULL,
	metadata_size bigint NOT NULL,
	PRIMARY KEY ( bucket_name, project_id, interval_start )
);
CREATE TABLE bucket_usages (
	id bytea NOT NULL,
	bucket_id bytea NOT NULL,
	rollup_end_time timestamp with time zone NOT NULL,
	remote_stored_data bigint NOT NULL,
	inline_stored_data bigint NOT NULL,
	remote_segments integer NOT NULL,
	inline_segments integer NOT NULL,
	objects integer NOT NULL,
	metadata_size bigint NOT NULL,
	repair_egress bigint NOT NULL,
	get_egress bigint NOT NULL,
	audit_egress bigint NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE bwagreements (
	serialnum text NOT NULL,
	storage_node_id bytea NOT NULL,
	uplink_id bytea NOT NULL,
	action bigint NOT NULL,
	total bigint NOT NULL,
	created_at timestamp with time zone NOT NULL,
	expires_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( serialnum )
);
CREATE TABLE certRecords (
	publickey bytea NOT NULL,
	id bytea NOT NULL,
	update_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE injuredsegments (
	id bigserial NOT NULL,
	info bytea NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE irreparabledbs (
	segmentpath bytea NOT NULL,
	segmentdetail bytea NOT NULL,
	pieces_lost_count bigint NOT NULL,
	seg_damaged_unix_sec bigint NOT NULL,
	repair_attempt_count bigint NOT NULL,
	PRIMARY KEY ( segmentpath )
);
CREATE TABLE nodes (
	id bytea NOT NULL,
	address text NOT NULL,
	protocol integer NOT NULL,
	type integer NOT NULL,
	email text NOT NULL,
	wallet text NOT NULL,
	free_bandwidth bigint NOT NULL,
	free_disk bigint NOT NULL,
	latency_90 bigint NOT NULL,
	audit_success_count bigint NOT NULL,
	total_audit_count bigint NOT NULL,
	audit_success_ratio double precision NOT NULL,
	uptime_success_count bigint NOT NULL,
	total_uptime_count bigint NOT NULL,
	uptime_ratio double precision NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	last_contact_success timestamp with time zone NOT NULL,
	last_contact_failure timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE projects (
	id bytea NOT NULL,
	name text NOT NULL,
	description text NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE registration_tokens (
	secret bytea NOT NULL,
	owner_id bytea,
	project_limit integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( secret ),
	UNIQUE ( owner_id )
);
CREATE TABLE serial_numbers (
	id serial NOT NULL,
	serial_number bytea NOT NULL,
	bucket_id bytea NOT NULL,
	expires_at timestamp NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE storagenode_bandwidth_rollups (
	storagenode_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	interval_seconds integer NOT NULL,
	action integer NOT NULL,
	allocated bigint NOT NULL,
	settled bigint NOT NULL,
	PRIMARY KEY ( storagenode_id, interval_start, action )
);
CREATE TABLE storagenode_storage_tallies (
	storagenode_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	total bigint NOT NULL,
	PRIMARY KEY ( storagenode_id, interval_start )
);
CREATE TABLE users (
	id bytea NOT NULL,
	full_name text NOT NULL,
	short_name text,
	email text NOT NULL,
	password_hash bytea NOT NULL,
	status integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE api_keys (
	id bytea NOT NULL,
	project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	key bytea NOT NULL,
	name text NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id ),
	UNIQUE ( key ),
	UNIQUE ( name, project_id )
);
CREATE TABLE project_members (
	member_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
	project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( member_id, project_id )
);
CREATE TABLE used_serials (
	serial_number_id integer NOT NULL REFERENCES serial_numbers( id ) ON DELETE CASCADE,
	storage_node_id bytea NOT NULL,
	PRIMARY KEY ( serial_number_id, storage_node_id )
);
CREATE INDEX bucket_name_project_id_interval_start_interval_seconds ON bucket_bandwidth_rollups ( bucket_name, project_id, interval_start, interval_seconds );
CREATE UNIQUE INDEX bucket_id_rollup ON bucket_usages ( bucket_id, rollup_end_time );
CREATE UNIQUE INDEX serial_number ON serial_numbers ( serial_number );
CREATE INDEX serial_numbers_expires_at_index ON serial_numbers ( expires_at );
CREATE INDEX storagenode_id_interval_start_interval_seconds ON storagenode_bandwidth_rollups ( storagenode_id, interval_start, interval_seconds );
