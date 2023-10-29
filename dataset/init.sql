CREATE TABLE public.logger (
	request varchar(255) NULL,
	response varchar(255) NULL,
	"method" varchar(10) NULL,
	code varchar(5) NULL,
	accesstime timestamp NULL
);