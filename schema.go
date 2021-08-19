package devicemngt

// DeviceManagementSchema ...
const DeviceManagementSchema = `
	CREATE TABLE if NOT EXISTS device_managements
	(
	    id text NOT NULL
	        CONSTRAINT device_managements_pkey
	        PRIMARY KEY,
	    device_id text NOT NULL UNIQUE,
			ip text NOT NULL,
			name text NOT NULL,
			platform text NOT NULL,
			os_name text NOT NULL,
			os_version text NOT NULL,
			app_version text NOT NULL,
			app_version_code text NOT NULL,
			browser_name text NOT NULL,
			browser_version text NOT NULL,
			auth_token text NOT NULL,
			fcm_token text,
			owner_id text NOT NULL,
			owner_type text NOT NULL,
	    first_sign_in_at timestamp with time zone not null,
			last_activity_at timestamp with time zone not null
	);
`
