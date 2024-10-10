package Functions

func CreateUserTable() (query string) {
	query = `
	DROP TABLE IF EXISTS public."user";
	CREATE TABLE public."user" (
    uuid uuid PRIMARY KEY,
    email character varying(100) UNIQUE,
    name character varying(100),
    lastname character varying(100),
    nickname character varying(100) UNIQUE,
    password text,
    tel character varying(13) UNIQUE,
    isverified boolean DEFAULT false
	);
	`
	return
}
func CreateCodeTable() (query string) {
	query = `
	DROP TABLE IF EXISTS public.codetbl;
	CREATE TABLE public.codetbl (
    id SERIAL PRIMARY KEY,
    code character varying(6),
    created_at timestamp without time zone DEFAULT now(),
    verifytype character varying(100)
	);
	`
	return
}
func CreateUuidEx() (query string) {
	query = `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    `
	return
}
func SignIn() (query string) {
	query = `
    CREATE OR REPLACE FUNCTION public.signin(_email character varying, _password character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
	if (SELECT COUNT(*) FROM public.user WHERE email=_email)>=1 then
      if (SELECT isverified FROM public.user WHERE email=_email AND password=_password)=true then
	  	RETURN json_build_object('statu', 'success','message','Giriş başarılı.');
	  else
	  	RETURN json_build_object('statu', 'error','message','Hesabınızı email veya telefon numarası kullanarak onaylatınız.');
	  end if;
	else
	  RETURN json_build_object('statu', 'error','message','Bu bilgilere sahip kullanıcı bulunmamaktadır.');
	end if; 
END;$$;
    `
	return
}
func SignUp() (query string) {
	query = `
    CREATE OR REPLACE FUNCTION public.signup(_email character varying, _name character varying, _lastname character varying, _nickname character varying, _password character varying, _tel character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
	IF (SELECT COUNT(*) FROM public.user WHERE email=_email)>=1 then
		RETURN json_build_object('statu', 'error','message','Bu email adresi başka kullanıcı tarafından alınmış.');
	ELSE
		IF (SELECT COUNT(*) FROM public.user WHERE nickname=_nickname)>=1 then
				    RETURN json_build_object('statu', 'error','message','Bu nickname başka kullanıcı tarafından alınmış.');
		ELSE 
			IF (SELECT COUNT(*) FROM public.user WHERE tel=_tel)>=1 then
				    RETURN json_build_object('statu', 'error','message','Bu telefon numarası başka kullanıcı tarafından alınmış.');
			ELSE 
		      INSERT INTO public.user (uuid,email,name,lastname,nickname,password,tel) VALUES (uuid_generate_v4(),_email,_name,_lastname,_nickname,_password,_tel);
			  RETURN json_build_object('statu', 'success','message','Kullanıcı başarıyla eklendi.');
			END IF;
		END IF;
	END IF;
END;$$;
    `
	return
}
func IsVerifiedAccount() (query string) {
	query = `
    CREATE OR REPLACE FUNCTION public.isaccountverified(_email character varying, _tel character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
	if (SELECT COUNT(*) FROM public.user WHERE email=_email OR tel=_tel)>0 then
		if (SELECT isverified FROM public.user WHERE email=_email OR tel=_tel)=true then
			RETURN json_build_object('statu', 'success','message','Kullanıcı hesabını zaten onaylamış.');
	 	else
			RETURN json_build_object('statu', 'error','message','Kullanıcının hesabı onaylı değil.');
		end if;
	 else
		RETURN json_build_object('statu', 'error','message','Böyle bir kullanıcı bilgisi mecvut değildir.');
	 end if;
END;$$;
    `
	return
}
func SetCode() (query string) {
	query = `
    CREATE OR REPLACE PROCEDURE public.setcode(IN _code character varying, IN _type character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
	TRUNCATE TABLE public.codetbl;
	INSERT INTO public.codetbl (code,verifytype) VALUES (_code,_type);
END;$$;
    `
	return
}
func VerifyAccount() (query string) {
	query = `
    CREATE OR REPLACE FUNCTION public.verifyaccount(_ocode character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
DECLARE 
    _code VARCHAR(6);
    _type VARCHAR(100);
    _latest_timestamp TIMESTAMPTZ;
BEGIN
    SELECT created_at INTO _latest_timestamp
    FROM public.codetbl
    ORDER BY id DESC
    LIMIT 1;

    IF _latest_timestamp IS NOT NULL AND CURRENT_TIMESTAMP - _latest_timestamp < INTERVAL '3 minutes' THEN
        SELECT code, verifytype INTO _code, _type
        FROM public.codetbl
        WHERE created_at = _latest_timestamp
        ORDER BY id DESC
        LIMIT 1;
        IF _code = _ocode THEN
            UPDATE public.user 
            SET isverified = true 
            WHERE email = _type OR tel = _type;

            RETURN json_build_object('status', 'success', 'message', 'Hesabınız onaylandı.');
        ELSE
            RETURN json_build_object('status', 'error', 'message', 'Girdiğiniz kod yanlış.');
        END IF;
    ELSE
        RETURN json_build_object('status', 'error', 'message', 'Zaman aşımına uğradı.');
    END IF;
END;$$;
    `
	return
}
