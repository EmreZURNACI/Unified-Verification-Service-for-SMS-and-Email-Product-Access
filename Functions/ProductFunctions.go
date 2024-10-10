package Functions

func CreateProductTable() (query string) {
	query = `
	DROP TABLE IF EXISTS public.product;
	CREATE TABLE public.product (
    id SERIAL PRIMARY KEY,
    marka character varying(255),
    model character varying(255) UNIQUE,
    isletimsistemi character varying(255)
	);
	`
	return
}
func Createproduct() (query string) {
	query = `
	CREATE OR REPLACE FUNCTION public.createproduct(_marka character varying, _model character varying, _isletimsistemi character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (SELECT COUNT(*) FROM public.product WHERE model = _model) > 0 THEN
        RETURN json_build_object('statu', 'error','message','Bu bilgilere sahip bir kayıt bulunmaktadır.');
    ELSE
        INSERT INTO public.product (marka, model, isletimsistemi) 
        VALUES (_marka, _model, _isletimsistemi);
        RETURN json_build_object('statu', 'success','message','Kayıt eklendi.');
    END IF;
END;$$;
	`
	return
}
func Deleteproduct() (query string) {
	query = `
	CREATE OR REPLACE FUNCTION public.deleteproduct(_id integer) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
	if (SELECT COUNT(*) FROM public.product WHERE id=_id)>0 then
		DELETE FROM public.product WHERE id=_id;
	    RETURN json_build_object('statu', 'success','message','Ürün silinmiştir');
	else 
	    RETURN json_build_object('statu', 'error','message','Bu id değerine sahip kayıt yoktur.');
	end if;
END;$$;
	`
	return
}
func Updateproduct() (query string) {
	query = `
	CREATE OR REPLACE FUNCTION public.updateproduct(_id integer, _marka character varying, _model character varying, _isletimsistemi character varying) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
	if (SELECT COUNT(*) FROM public.product WHERE id=_id)=0 then
		RETURN json_build_object('statu', 'error','message','Bu id değerine sahip kayıt yok.');
	elsif (SELECT COUNT(*) FROM public.product WHERE model=_model)>=1 then
		RETURN json_build_object('statu', 'error','message','Bu özelliklere sahip zaten ürün bulunmaktadır.');
	else
		UPDATE public.product SET marka=_marka,model=_model,isletimsistemi=_isletimsistemi WHERE id=_id;
		RETURN json_build_object('statu', 'success','message','kayıt güncellendi.');
	end if;
END;$$;
	`
	return
}
func InsertFakeData() (query string) {
	query = `
	CREATE OR REPLACE FUNCTION public.InsertFakeData() RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    sayac INT := 1;
BEGIN
    WHILE sayac <= 1000 LOOP
        INSERT INTO product(marka, model, isletimsistemi)
        VALUES ('Marka_' || sayac, 'Model_' || sayac, 'System_' || sayac);
        sayac := sayac + 1;
    END LOOP;
END;$$;
SELECT public.InsertFakeData();
	`
	return
}
