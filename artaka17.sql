--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Ubuntu 13.0-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.0 (Ubuntu 13.0-1.pgdg20.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: admins; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.admins (
    id bigint NOT NULL,
    phone character varying(100),
    username character varying(255) NOT NULL,
    create_dtm timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    email character varying(100) NOT NULL,
    secret_password text
);


ALTER TABLE public.admins OWNER TO postgres;

--
-- Name: admins_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.admins_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admins_id_seq OWNER TO postgres;

--
-- Name: admins_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.admins_id_seq OWNED BY public.admins.id;


--
-- Name: onlinesales; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.onlinesales (
    id integer,
    create_dtm timestamp without time zone,
    sales_id character varying(50),
    user_id character varying(50),
    outlet_id character varying(50),
    customer_id character varying(50),
    customer json,
    products json,
    subtotal integer,
    total_diskon integer,
    total_tax json,
    total_bill integer,
    payment_method character varying(50),
    payment_account character varying(50),
    payment_due_date character varying(50),
    total_payment integer,
    expedition character varying(50),
    service character varying(50),
    weight integer,
    delivery_cost integer,
    notes character varying(100),
    total_buy_cost integer,
    payment_date character varying(50),
    reward_id character varying(50),
    points_redeem integer,
    order_status character varying(50),
    shipment_number character varying(50)
);


ALTER TABLE public.onlinesales OWNER TO postgres;

--
-- Name: outlets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.outlets (
    id integer,
    create_dtm timestamp without time zone,
    user_id character varying(50),
    outlet_id character varying(50),
    nama character varying(100),
    address character varying(300),
    phone character varying(30),
    business_category character varying(50),
    is_active character varying(3),
    accounts json,
    images json,
    mini_website_url character varying(128),
    is_online_store_active character varying(3)
);


ALTER TABLE public.outlets OWNER TO postgres;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    phone character varying(255) NOT NULL,
    content text NOT NULL,
    "boolean" text DEFAULT '0'::text,
    author_id bigint NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO postgres;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: sales; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales (
    id integer,
    create_dtm timestamp without time zone,
    sales_id character varying(50),
    user_id character varying(50),
    outlet_id character varying(50),
    sales_type character varying(50),
    customer_id character varying(50),
    products json,
    subtotal integer,
    total_diskon integer,
    total_bill integer,
    payment_method character varying(50),
    payment_due_date character varying(50),
    total_payment integer,
    exchange integer,
    notes character varying(100),
    total_buy_cost integer,
    payment_date character varying(20),
    total_tax json,
    reward_id character varying(50),
    points_redeem integer
);


ALTER TABLE public.sales OWNER TO postgres;

--
-- Name: saved_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.saved_orders (
    id integer,
    create_dtm timestamp without time zone,
    user_id character varying(50),
    outlet_id character varying(50),
    saved_orders_id character varying(50),
    name character varying(50),
    phone character varying(20),
    orders json,
    table_id character varying(20)
);


ALTER TABLE public.saved_orders OWNER TO postgres;

--
-- Name: subscribers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subscribers (
    id integer,
    create_dtm timestamp without time zone,
    user_id character varying(50),
    email character varying(100),
    owner_name character varying(100),
    secret_password character varying(50),
    fcm_token character varying(200),
    idcard_name character varying(50),
    idcard_number character varying(50),
    bank_holder_name character varying(256),
    bank_name character varying(256),
    bank_account character varying(256),
    idcard_image json,
    referral_code character varying(50)
);


ALTER TABLE public.subscribers OWNER TO postgres;

--
-- Name: admins id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins ALTER COLUMN id SET DEFAULT nextval('public.admins_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.admins (id, phone, username, create_dtm, email, secret_password) FROM stdin;
\.


--
-- Data for Name: onlinesales; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.onlinesales (id, create_dtm, sales_id, user_id, outlet_id, customer_id, customer, products, subtotal, total_diskon, total_tax, total_bill, payment_method, payment_account, payment_due_date, total_payment, expedition, service, weight, delivery_cost, notes, total_buy_cost, payment_date, reward_id, points_redeem, order_status, shipment_number) FROM stdin;
53	2020-08-06 20:54:00.454307	OS-20200806-0000001	+6281311666268	OTL-001		{"name":"Elina","handphone":"0811109215","email":"elina@gmail.com","alamat":"Kota Wisataa","kecamatan":"Gunung Putri, Bogor, Jawa Barat","kodepos":"16911","subdistrics_info":{"subdistrict_id":"1039","city_id":"78","province_id":"9","subdistrict_name":"Gunung Putri","city":"Bogor","type":"Kabupaten","province":"Jawa Barat","postal_code":"16911"},"fcm_web_token":"eN8gWkiI_KpBOPIIWJugtW:APA91bHU9HwQoIPGP_4PNzJTEdWem53uWqA5uYC18h-cSZK140B3yxZKr3FRsLx8Z8O9qswiHOPVbnVT6Jmy-Q-qpwhESoG7aD_NqA4mBGkn2wcgbfKNrhl0Q3uQDBhbOK9rwKpDSDpX"}	[{"sku_id":"FP-0001","name":"Nasi Goreng","category":"Manual","variant":"","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":40000,"buy_cost_discounted":0,"sell_cost":50000,"weight":0,"units":"Pieces","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000	Cash On Delivery (COD)			50000	Diantar Oleh Penjual	Rp 0	0	0		0			0	Baru	
54	2020-08-06 20:54:52.471212	OS-20200806-0000002	+6281311666268	OTL-001		{"name":"Elina","handphone":"0811109215","email":"elina@gmail.com","alamat":"Kota Wisataa","kecamatan":"Gunung Putri, Bogor, Jawa Barat","kodepos":"16911","subdistrics_info":{"subdistrict_id":"1039","city_id":"78","province_id":"9","subdistrict_name":"Gunung Putri","city":"Bogor","type":"Kabupaten","province":"Jawa Barat","postal_code":"16911"},"fcm_web_token":"eN8gWkiI_KpBOPIIWJugtW:APA91bHU9HwQoIPGP_4PNzJTEdWem53uWqA5uYC18h-cSZK140B3yxZKr3FRsLx8Z8O9qswiHOPVbnVT6Jmy-Q-qpwhESoG7aD_NqA4mBGkn2wcgbfKNrhl0Q3uQDBhbOK9rwKpDSDpX"}	[{"sku_id":"FP-0001","name":"Nasi Goreng","category":"Manual","variant":"","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":40000,"buy_cost_discounted":0,"sell_cost":50000,"weight":0,"units":"Pieces","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000	BNI Virtual Account	8808839462902966		50000	Diantar Oleh Penjual	Gratis	0	0		0			0	Baru	
55	2020-08-06 20:56:33.283191	OS-20200806-0000003	+6281311666268	OTL-001		{"name":"Bembs","handphone":"0811500500","email":"Bembs@gmail.com","alamat":"kota wisata","kecamatan":"Mampang Prapatan, Jakarta Selatan, DKI Jakarta","kodepos":"12230","subdistrics_info":{"subdistrict_id":"2107","city_id":"153","province_id":"6","subdistrict_name":"Mampang Prapatan","city":"Jakarta Selatan","type":"Kota","province":"DKI Jakarta","postal_code":"12230"},"fcm_web_token":"eMPugwrNx_07VfydQmb9XA:APA91bH0UrMyVODKCjo3J0hnUxvLss7b2awH3y7OkZZex5dcyG1-sqOpuOMuJRgemFJf6SHZthTzeSN30RQj8HfI6LcJsLMEaYhcHm3Gq-n8vKL_RjesRamtbB7xSEEnIeN-cfzuqGZ2"}	[{"sku_id":"FP-0001","name":"Nasi Goreng","category":"Manual","variant":"","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":40000,"buy_cost_discounted":0,"sell_cost":50000,"weight":0,"units":"Pieces","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000	BCA Virtual Account			50000	Diantar Oleh Penjual	Gratis	0	0		0			0	Baru	
56	2020-08-06 21:17:48.907573	OS-20200806-0000004	+6281311666268	OTL-001		{"name":"Fajar","handphone":"081311666268","email":"mohfajar.173@gmail.com","alamat":"Jalan prof moch yamin no 3 menteng jakarta pusat","kecamatan":"Menteng, Jakarta Pusat, DKI Jakarta","kodepos":"10310","subdistrics_info":{"subdistrict_id":"2099","city_id":"152","province_id":"6","subdistrict_name":"Menteng","city":"Jakarta Pusat","type":"Kota","province":"DKI Jakarta","postal_code":"10540"},"fcm_web_token":"fqcLwsQ0Lo0XXA7sVrQW5Z:APA91bGyOTf8Zn_l_WU2D6gwOm4Nl2TzbaEzydjyEdJ4TDf5zmn_9vWjhy8bBVR0m3INRF4r7C4hNGUO2J7fGIcEy6MSCxktkARBhWnDYp5JuLvCAND1r_lHpfqwhk6PUN0HLtJXZU4F"}	[{"sku_id":"FP-0001","name":"Nasi Goreng","category":"Manual","variant":"","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":40000,"buy_cost_discounted":0,"sell_cost":50000,"weight":0,"units":"Pieces","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000	BCA Virtual Account			50000	Diantar Oleh Penjual	Gratis	0	0	tanpa catatan	0			0	Baru	
57	2020-08-06 22:20:41.305311	OS-20200806-0000005	+6281311666268	OTL-001		{"name":"Hasan","handphone":"0811987905","email":"Hasan@gmail.com","alamat":"Bintato","kecamatan":"Jombang, Cilegon, Banten","kodepos":"42417","subdistrics_info":{"subdistrict_id":"1466","city_id":"106","province_id":"3","subdistrict_name":"Jombang","city":"Cilegon","type":"Kota","province":"Banten","postal_code":"42417"},"fcm_web_token":"fj_kSjFFsNoTcNKJPB4gPh:APA91bFtbx9zuQVHjfIfheVvGFdbHH5T8uqgVuwsgiiuTrgy50QQ1gLCnb3rAmtYpLA4RBO3hVCLCaoMrGuDIOubGoI3Zv-_qfTn_ynxR3Yi-7qQI9MUjeMDrr5w-vntWZblW1yDSXB7"}	[{"sku_id":"FP-0001","name":"Nasi Goreng","category":"Manual","variant":"","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":40000,"buy_cost_discounted":0,"sell_cost":50000,"weight":0,"units":"Pieces","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000	Cash On Delivery (COD)			50000	Diantar Oleh Penjual	Gratis	0	0		0			0	Baru	
293	2020-09-22 23:33:18.146481	OS-20200922-0000002	+6282264291947	OTL-001		{"name":"moe","handphone":"087774655211","email":"modhenk11@gmail.com","alamat":"panji","kecamatan":"Panji, Situbondo, Jawa Timur","kodepos":"68316","subdistrics_info":{"subdistrict_id":"5775","city_id":"418","province_id":"11","subdistrict_name":"Panji","city":"Situbondo","type":"Kabupaten","province":"Jawa Timur","postal_code":"68316"},"fcm_web_token":""}	[{"sku_id":"FP-0010","name":"Mister Moron","category":"Lainnya","variant":"Nic3","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":140000,"buy_cost_discounted":0,"sell_cost":160000,"weight":0,"units":"Botol","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	160000	0	{"PPN":0,"Service Charge":0,"PB1":0}	160000	Cash On Delivery (COD)			160000	Diantar Oleh Penjual	Gratis	0	0		0			0	Order Selesai	
292	2020-09-22 23:22:05.1295	OS-20200922-0000001	+6282264291947	OTL-001		{"name":"moe","handphone":"085230636631","email":"modhenk11@gmail.com","alamat":"panji","kecamatan":"Panji, Situbondo, Jawa Timur","kodepos":"68316","subdistrics_info":{"subdistrict_id":"5775","city_id":"418","province_id":"11","subdistrict_name":"Panji","city":"Situbondo","type":"Kabupaten","province":"Jawa Timur","postal_code":"68316"},"fcm_web_token":""}	[{"sku_id":"FP-0001","name":"WIRE","category":"Lainnya","variant":"baby alien","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":15000,"buy_cost_discounted":15000,"sell_cost":30000,"weight":2,"units":"Botol","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	30000	0	{"PPN":0,"Service Charge":0,"PB1":0}	30000	Cash On Delivery (COD)			30000	Diantar Oleh Penjual	Gratis	2	0		15000			0	Order Selesai	
315	2020-09-24 16:49:04.845486	#20092416470081	+6282264291947	OTL-001	CUS-0283	{"name":"B. M MELON SUSU","handphone":"6285336256525","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2875","create_dtm":"2020-09-22T22:57:03.624085Z","sku_id":"FP-0008","user_id":"+6282264291947","outlet_id":"OTL-001","name":"BOMBER","category":"Lainnya","variant":"OB Strawberry","units":"Botol","weight":5,"quantity":4,"minimum_quantity":0,"description":"","buy_cost":28000,"sell_cost":45000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/95948069-9c65-45d4-9bc1-27329b5f3c11.jpg?alt=media&token=c4c24008-a51f-4517-8d62-95120c506fe4"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":28000,"is_active":"Yes","salestype_up":0}]	45000	0	{"PPN":0,"Service Charge":0,"PB1":0}	45000	Transfer Bank			45000	Diantar Oleh Penjual		5	0		28000			0	Order Selesai	
316	2020-09-24 16:57:13.972554	#20092416502324	+6282264291947	OTL-001	CUS-0361	{"name":"ngeboel banabeast4","handphone":"6282246624922","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2883","create_dtm":"2020-09-22T23:06:56.826051Z","sku_id":"FP-0016","user_id":"+6282264291947","outlet_id":"OTL-001","name":"COTTON","category":"Lainnya","variant":"KENDO","units":"Pack","weight":0,"quantity":2,"minimum_quantity":0,"description":"","buy_cost":26000,"sell_cost":40000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/40313b69-a0a5-4536-a629-35bdfa8e3295.jpg?alt=media&token=a00c85bc-eb0b-4074-9186-9b76c09b70bd"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":12,"outlets":["OTL-001"],"buy_cost_discounted":26000,"is_active":"Yes","salestype_up":0}]	40000	0	{"PPN":0,"Service Charge":0,"PB1":0}	40000	Transfer Bank			40000	Diantar Oleh Penjual		0	0		26000			0	Order Selesai	
317	2020-09-24 19:00:03.338767	#20092323322023	+6282264291947	OTL-001	CUS-0422	{"name":"kejar tayang","handphone":"6282330316008","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2868","create_dtm":"2020-09-22T22:30:35.793763Z","sku_id":"FP-0001","user_id":"+6282264291947","outlet_id":"OTL-001","name":"WIRE","category":"Lainnya","variant":"baby alien","units":"Pieces","weight":2,"quantity":3,"minimum_quantity":0,"description":"harga per pasang ya dual 0,18-0,24ohm","buy_cost":15000,"sell_cost":30000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/82d893bb-ef5a-4d2e-ad81-e05e78fcb439.jpg?alt=media&token=685d2fda-4676-47fa-8636-dfdabd111bfe"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":2,"outlets":["OTL-001"],"buy_cost_discounted":15000,"is_active":"Yes","salestype_up":0},{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2873","create_dtm":"2020-09-22T22:57:03.619729Z","sku_id":"FP-0006","user_id":"+6282264291947","outlet_id":"OTL-001","name":"BOMBER","category":"Lainnya","variant":"melon susu","units":"Botol","weight":5,"quantity":3,"minimum_quantity":0,"description":"","buy_cost":25000,"sell_cost":45000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/f6f0260c-02fa-4355-bf5e-a841618ebf27.jpg?alt=media&token=041779f5-2f06-4eee-bee3-da00a62465bc"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":25000,"is_active":"Yes","salestype_up":0}]	75000	0	{"PPN":0,"Service Charge":0,"PB1":0}	75000	Transfer Bank			75000	Diantar Oleh Penjual		7	0		40000			0	Order Selesai	
319	2020-09-24 20:48:29.035866	#20092420460214	+6282264291947	OTL-001	CUS-0102	{"name":"maman","handphone":"6289682198186","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2883","create_dtm":"2020-09-22T23:06:56.826051Z","sku_id":"FP-0016","user_id":"+6282264291947","outlet_id":"OTL-001","name":"COTTON","category":"Lainnya","variant":"KENDO","units":"Pack","weight":0,"minimum_quantity":0,"description":"","buy_cost":26000,"sell_cost":40000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/40313b69-a0a5-4536-a629-35bdfa8e3295.jpg?alt=media&token=a00c85bc-eb0b-4074-9186-9b76c09b70bd"],"rawmaterial":[],"is_stock_tracked":"Yes","outlets":["OTL-001"],"buy_cost_discounted":26000,"is_active":"Yes","salestype_up":0,"quantity":1,"number_sold":15}]	40000	0	{"PPN":0,"Service Charge":0,"PB1":0}	40000	Transfer Bank			40000	Diantar Oleh Penjual		0	0		26000			0	Order Selesai	
360	2020-09-28 18:11:16.649354	#20092817530866	+6282264291947	OTL-001	CUS-0408	{"name":"ngeboel banabeast bri","handphone":"6281290447519","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3075","create_dtm":"2020-09-24T21:51:42.613226Z","sku_id":"FP-0016","user_id":"+6282264291947","outlet_id":"OTL-001","name":"COTTON","category":"Lainnya","variant":"Beck-On","units":"Pack","weight":10,"quantity":1,"minimum_quantity":0,"description":"","buy_cost":8000,"sell_cost":20000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/7f383cf4-0bb2-47d0-99e6-10bf898180c3.jpg?alt=media&token=f5b27104-6710-4a85-8b43-456ac6cf10fa"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":8000,"is_active":"Yes","salestype_up":0}]	20000	0	{"PPN":0,"Service Charge":0,"PB1":0}	20000	Transfer Bank			20000	Diantar Oleh Penjual		10	0		8000			0	Order Selesai	
365	2020-09-29 14:38:29.80779	#20092914375619	+6282264291947	OTL-001	CUS-0383	{"name":"alacarte prjekan","handphone":"6285204862177","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2868","create_dtm":"2020-09-22T22:30:35.793763Z","sku_id":"FP-0001","user_id":"+6282264291947","outlet_id":"OTL-001","name":"WIRE","category":"Lainnya","variant":"baby alien","units":"Pieces","weight":2,"quantity":14,"minimum_quantity":0,"description":"harga per pasang ya dual 0,18-0,24ohm","buy_cost":15000,"sell_cost":30000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/11ac7067-637e-48a3-a679-2dc4112026b2.jpg?alt=media&token=87021365-c5cd-4791-9081-5e89cb7dd7e8"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":3,"outlets":["OTL-001"],"buy_cost_discounted":15000,"is_active":"Yes","salestype_up":0}]	30000	0	{"PPN":0,"Service Charge":0,"PB1":0}	30000	Transfer Bank			30000	Diantar Oleh Penjual		2	0		15000			0	Order Selesai	
366	2020-09-29 18:56:52.182456	#20092918290567	+6282264291947	OTL-001	CUS-0329	{"name":"B M CHAENG","handphone":"6289631007240","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2875","create_dtm":"2020-09-22T22:57:03.624085Z","sku_id":"FP-0008","user_id":"+6282264291947","outlet_id":"OTL-001","name":"BOMBER","category":"Lainnya","variant":"OB Strawberry","units":"Botol","weight":5,"quantity":3,"minimum_quantity":0,"description":"","buy_cost":28000,"sell_cost":45000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/95948069-9c65-45d4-9bc1-27329b5f3c11.jpg?alt=media&token=c4c24008-a51f-4517-8d62-95120c506fe4"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":1,"outlets":["OTL-001"],"buy_cost_discounted":28000,"is_active":"Yes","salestype_up":0}]	45000	0	{"PPN":0,"Service Charge":0,"PB1":0}	45000	Transfer Bank			45000	Diantar Oleh Penjual		5	0		28000			0	Order Selesai	
386	2020-09-30 22:34:43.389576	#20092316392033	+6282264291947	OTL-001	CUS-0028	{"name":"pringgo","handphone":"6281330756067","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2910","create_dtm":"2020-09-23T12:50:56.414128Z","sku_id":"FP-0019","user_id":"+6282264291947","outlet_id":"OTL-001","name":"KOPI TIRAMISU","category":"Lainnya","variant":"Kopi tiramisu nic3","units":"Botol","weight":6,"quantity":1,"minimum_quantity":0,"description":"","buy_cost":55000,"sell_cost":65000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/d48ce5ca-ae73-49ce-b3d2-584bf76c3352.jpg?alt=media&token=d9589693-e000-4d3d-82e2-12cf9104269f"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":2,"outlets":["OTL-001"],"buy_cost_discounted":0,"is_active":"Yes","salestype_up":0},{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2869","create_dtm":"2020-09-22T22:30:35.797025Z","sku_id":"FP-0002","user_id":"+6282264291947","outlet_id":"OTL-001","name":"WIRE","category":"Lainnya","variant":"mini alien","units":"Pieces","weight":2,"quantity":1,"minimum_quantity":0,"description":"harga per pasang ya dual 0,17-0,19 ohm","buy_cost":20000,"sell_cost":30000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/0d207d0f-f80c-4816-b315-989e19334d9c.jpg?alt=media&token=f713ade6-134e-4f87-8cd6-32376f4699b2"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":20000,"is_active":"Yes","salestype_up":0}]	95000	0	{"PPN":0,"Service Charge":0,"PB1":0}	100000	Transfer Bank			100000	Ojek Online		8	5000		20000			0	Order Selesai	
419	2020-10-02 11:11:14.208947	OS-20201002-0000001	+6282264291947	OTL-001		{"name":"Bagus Fahroni Aziiz","handphone":"085336256525","email":"bagus.cool49@gmail.com","alamat":"Prajekan por RT.01 RW.05","kecamatan":"Prajekan, Bondowoso, Jawa Timur","kodepos":"68219","subdistrics_info":{"subdistrict_id":"1159","city_id":"86","province_id":"11","subdistrict_name":"Prajekan","city":"Bondowoso","type":"Kabupaten","province":"Jawa Timur","postal_code":"68219"},"fcm_web_token":""}	[{"sku_id":"FP-0004","name":"BOMBER","category":"Lainnya","variant":"strawberry vanilla","modifiers_price":0,"modifiers_option":"","number_orders":1,"buy_cost":25000,"buy_cost_discounted":25000,"sell_cost":45000,"weight":5,"units":"Botol","salestype_up":0,"discount_info":{"name":"","amount":0},"taxInfo":[],"description":""}]	45000	0	{"PPN":0,"Service Charge":0,"PB1":0}	45000	Cash On Delivery (COD)			45000	Diantar Oleh Penjual	Gratis	5	0	COD kantor PLN situbondo	25000			0	Order Selesai	
423	2020-10-02 13:04:18.332641	#20100212084759	+6282264291947	OTL-001	CUS-0432	{"name":"icak vape besuki","handphone":"6281233883244","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":40000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3078","create_dtm":"2020-09-24T21:56:22.110913Z","sku_id":"FP-0017","user_id":"+6282264291947","outlet_id":"OTL-001","name":"KENDO VAPE COTTON ","category":"Lainnya","variant":"","units":"Pack","weight":10,"quantity":7,"minimum_quantity":0,"description":"","buy_cost":26000,"sell_cost":40000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/68a03aa2-836d-44d8-94fb-6522aef59f70.jpg?alt=media&token=a8bc27f3-9ff4-4757-b9bd-2c0ffad2b281"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":26000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0},{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":30000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3449","create_dtm":"2020-10-01T19:38:53.414475Z","sku_id":"FP-0022","user_id":"+6282264291947","outlet_id":"OTL-001","name":"glass RTA ZEUS X /SULTAN X","category":"Lainnya","variant":"","units":"Butir","weight":0,"quantity":3,"minimum_quantity":0,"description":"yg cembung ukuran 5ml","buy_cost":12000,"sell_cost":30000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/94f2a15e-8a4c-409f-a3e3-9308518ee84b.jpg?alt=media&token=fb29242e-0c66-4d17-ad12-67f07c03a291"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":1,"outlets":["OTL-001"],"buy_cost_discounted":12000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	70000	0	{"PPN":0,"Service Charge":0,"PB1":0}	70000				70000	Diantar Oleh Penjual		10	0		38000			0	Order Selesai	
424	2020-10-02 14:07:31.7114	#20100212200222	+6282264291947	OTL-001	CUS-0423	{"name":"banabeast patokan5","handphone":"6287712994897","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":50000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2882","create_dtm":"2020-09-22T23:04:53.893061Z","sku_id":"FP-0015","user_id":"+6282264291947","outlet_id":"OTL-001","name":"CHARGER AWT","category":"Lainnya","variant":"","units":"Pack","weight":1,"quantity":2,"minimum_quantity":0,"description":"","buy_cost":35000,"sell_cost":50000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/ed2a85e4-3814-45f1-b615-8df1e5e86feb.jpg?alt=media&token=6f049cfc-c6e2-43b4-8aa8-f3fe16971241"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":3,"outlets":["OTL-001"],"buy_cost_discounted":35000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	50000	0	{"PPN":0,"Service Charge":0,"PB1":0}	50000				50000	Diantar Oleh Penjual		1	0		35000			0	Order Selesai	
427	2020-10-02 16:09:07.246859	#20100214302549	+6282264291947	OTL-001	CUS-0413	{"name":"baby alien gldk macan","handphone":"6282331581769","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":160000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2878","create_dtm":"2020-09-22T22:59:44.294888Z","sku_id":"FP-0011","user_id":"+6282264291947","outlet_id":"OTL-001","name":"Mister Moron","category":"Lainnya","variant":"Nic6","units":"Botol","weight":0,"quantity":2,"minimum_quantity":0,"description":"","buy_cost":140000,"sell_cost":160000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/6461f18d-b46f-4a5c-9238-09866a7f2866.jpg?alt=media&token=4a27f712-203d-413b-b80b-ebc5d595db5d"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":1,"outlets":["OTL-001"],"buy_cost_discounted":140000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	160000	0	{"PPN":0,"Service Charge":0,"PB1":0}	160000				160000	Diantar Oleh Penjual		0	0		140000			0	Order Selesai	
428	2020-10-02 16:09:20.706769	#20100108504730	+6282264291947	OTL-001	CUS-0413	{"name":"baby alien gldk macan","handphone":"6282331581769","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3355","create_dtm":"2020-10-01T08:48:13.886606Z","sku_id":"FP-0020","user_id":"+6282264291947","outlet_id":"OTL-001","name":"TOKYO GREEN BEAN","category":"Lainnya","variant":"","units":"Botol","weight":60,"quantity":1,"minimum_quantity":0,"description":"","buy_cost":127500,"sell_cost":150000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/e0a7968c-404a-4674-84dc-ccad48a2947c.jpg?alt=media&token=60f7d40d-6c04-4fdb-94ec-fa54baaaee82"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":127500,"is_active":"Yes","salestype_up":0}]	150000	0	{"PPN":0,"Service Charge":0,"PB1":0}	150000				150000	Diantar Oleh Penjual		60	0		127500			0	Order Selesai	
464	2020-10-03 19:43:07.207915	#20100310225586	+6282264291947	OTL-001	CUS-0434	{"name":"B.M KUMPUL2","handphone":"6281259124305","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":45000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"2873","create_dtm":"2020-09-22T22:57:03.619729Z","sku_id":"FP-0006","user_id":"+6282264291947","outlet_id":"OTL-001","name":"BOMBER","category":"Lainnya","variant":"melon susu","units":"Botol","weight":5,"quantity":1,"minimum_quantity":0,"description":"","buy_cost":25000,"sell_cost":45000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/f6f0260c-02fa-4355-bf5e-a841618ebf27.jpg?alt=media&token=041779f5-2f06-4eee-bee3-da00a62465bc"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":5,"outlets":["OTL-001"],"buy_cost_discounted":25000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	45000	0	{"PPN":0,"Service Charge":0,"PB1":0}	45000				45000	Diantar Oleh Penjual		5	0		25000			0	Order Selesai	
465	2020-10-03 19:47:55.461218	#20100317005877	+6282264291947	OTL-001	CUS-0293	{"name":"coil fused 2","handphone":"6281259556561","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":40000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3078","create_dtm":"2020-09-24T21:56:22.110913Z","sku_id":"FP-0017","user_id":"+6282264291947","outlet_id":"OTL-001","name":"KENDO VAPE COTTON ","category":"Lainnya","variant":"","units":"Pack","weight":10,"quantity":5,"minimum_quantity":0,"description":"","buy_cost":26000,"sell_cost":40000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/68a03aa2-836d-44d8-94fb-6522aef59f70.jpg?alt=media&token=a8bc27f3-9ff4-4757-b9bd-2c0ffad2b281"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":2,"outlets":["OTL-001"],"buy_cost_discounted":26000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	40000	0	{"PPN":0,"Service Charge":0,"PB1":0}	40000				40000	Diantar Oleh Penjual		10	0		26000			0	Order Selesai	
494	2020-10-05 20:20:20.757676	#20100514483832	+6282264291947	OTL-001	CUS-0054	{"name":"arin","handphone":"6282298859637","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":65000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3919","create_dtm":"2020-10-05T14:46:07.874339Z","sku_id":"FP-0024","user_id":"+6282264291947","outlet_id":"OTL-001","name":"NGEBOEL banabeast nic3","category":"Lainnya","variant":"","units":"Botol","weight":0,"quantity":1,"minimum_quantity":0,"description":"","buy_cost":55000,"sell_cost":65000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/2755e604-9e8b-4418-8aaf-7158555d3a48.jpg?alt=media&token=959e31a2-4f5c-4955-b32e-a64928a33be0"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":55000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	65000	0	{"PPN":0,"Service Charge":0,"PB1":0}	65000				65000	Diantar Oleh Penjual		0	0		55000			0	Order Selesai	
495	2020-10-05 20:20:51.726301	#20100413332217	+6282264291947	OTL-001	CUS-0413	{"name":"baby alien gldk macan","handphone":"6282331581769","email":"","alamat":"","kecamatan":"","kodepos":"","subdistrics_info":"","fcm_web_token":""}	[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":300000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3853","create_dtm":"2020-10-04T13:30:44.161915Z","sku_id":"FP-0024","user_id":"+6282264291947","outlet_id":"OTL-001","name":"toolkit &tarterie","category":"Jasa","variant":"","units":"Botol","weight":0,"quantity":1,"minimum_quantity":0,"description":"pesanan boss singo 🤣","buy_cost":270000,"sell_cost":300000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/2029770d-3b39-4b6b-9ec1-11806ed13a35.jpg?alt=media&token=0b6208c6-4aad-4183-b17e-859ae180d005"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":270000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}]	300000	0	{"PPN":0,"Service Charge":0,"PB1":0}	300000				300000	Diantar Oleh Penjual		0	0		270000			0	Order Selesai	
\.


--
-- Data for Name: outlets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.outlets (id, create_dtm, user_id, outlet_id, nama, address, phone, business_category, is_active, accounts, images, mini_website_url, is_online_store_active) FROM stdin;
72	2020-08-14 13:47:20.314967	+628116261101	OTL-001	FORZA GABE	Cluster Alicante blok AB7 No 50|Pagedangan, Tangerang|Banten|6284|455|3|15914	+628116261101	Toko Retail	Yes	{"kas_bank":0,"aset":0,"piutang":0,"hutang":0,"accounting_start_date":"14-Aug-2020"}	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	https://orderin.id/forzagabe	Yes
65	2020-08-06 11:22:16.736338	+62811196196	OTL-001	CENTRALatte	Jalan Mokmer 1|Sawah Besar, Jakarta Pusat|DKI Jakarta|2100|152|6|10540	+62811196196	Makanan & Minuman	Yes	{"kas_bank":0,"aset":0,"piutang":0,"hutang":0,"accounting_start_date":"06-Aug-2020"}	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	https://orderin.id/centralatte	Yes
55	2020-07-26 17:24:47.505511	+6281311666268	OTL-001	Nibbls	Menteng|Menteng, Jakarta Pusat|DKI Jakarta|2099|152|6|10540	081311666268	Makanan & Minuman	Yes	{"kas_bank":0,"aset":0,"piutang":0,"hutang":0,"accounting_start_date":"26-Jul-2020"}	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	https://orderin.id/nibbls	Yes
226	2020-09-22 22:25:33.206303	+6282264291947	OTL-001	MODVAPE	SITUBONDO KOTA|Situbondo, Situbondo|Jawa Timur|5776|418|11|68316	+6282264291947	Toko dan Usaha lainnya	Yes	{"kas_bank":0,"aset":0,"piutang":0,"hutang":0,"accounting_start_date":"22-Sep-2020"}	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	https://orderin.id/modhenk	Yes
366	2020-10-05 08:16:01.045041	+6281932809265	OTL-001	Mitra Mas Bendol	Jalan makam|Pinang (Penang), Tangerang|Banten|6308|456|3|15111	+6281932809265	Toko dan Usaha lainnya	Yes	{"kas_bank":0,"aset":0,"piutang":0,"hutang":0,"accounting_start_date":"05-Oct-2020"}	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	https://orderin.id/hariscimol	Yes
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.posts (id, phone, content, "boolean", author_id, updated_at) FROM stdin;
\.


--
-- Data for Name: sales; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales (id, create_dtm, sales_id, user_id, outlet_id, sales_type, customer_id, products, subtotal, total_diskon, total_bill, payment_method, payment_due_date, total_payment, exchange, notes, total_buy_cost, payment_date, total_tax, reward_id, points_redeem) FROM stdin;
185	2020-07-26 18:57:09.759585	SL-20200726-0000001	+6281311666268	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"","name":"Nasi Goreng","category":"Manual Input","variant":"","modifier_option":"","number_orders":1,"buy_cost":40000,"sell_cost":50000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0}]	50000	0	50000	Tunai	Invalid Date	50000	0		40000		{"PPN":0,"Service Charge":0,"PB1":0}		0
430	2020-09-22 22:35:12.527945	SL-20200922-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0001","name":"WIRE","category":"Lainnya","variant":"baby alien","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":15000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Tunai	Invalid Date	30000	0		15000		{"PPN":0,"Service Charge":0,"PB1":0}		0
442	2020-09-23 11:59:01.603945	SL-20200923-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0017","name":"COTTON","category":"Lainnya","variant":"BECK.ON","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":8000,"sell_cost":20000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	20000	0	20000	Tunai	Invalid Date	20000	0		8000		{"PPN":0,"Service Charge":0,"PB1":0}		0
443	2020-09-23 12:01:00.835918	SL-20200923-0000002	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0018","name":"NGEBOEL","category":"Lainnya","variant":"Kopi tiramisu nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Tunai	Invalid Date	65000	0		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
444	2020-09-23 12:51:38.502935	SL-20200923-0000003	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0019","name":"KOPI TIRAMISU","category":"Lainnya","variant":"Kopi tiramisu nic3","modifier_option":"","modifier_price":0,"number_orders":2,"buy_cost":0,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	130000	0	130000	Tunai	Invalid Date	130000	0		0		{"PPN":0,"Service Charge":0,"PB1":0}		0
447	2020-09-23 16:30:45.633518	SL-20200923-0000004	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0004","name":"BOMBER","category":"Lainnya","variant":"strawberry vanilla","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""},{"sku_id":"FP-0017","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	110000	0	110000	Uang Pas	Invalid Date	110000	0		80000		{"PPN":0,"Service Charge":0,"PB1":0}		0
449	2020-09-23 18:13:36.975194	SL-20200923-0000005	+6282264291947	OTL-001	Bawa Pulang	CUS-0243	[{"sku_id":"FP-0009","name":"BOMBER","category":"Lainnya","variant":"OB Blueberry","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":28000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Uang Pas	Invalid Date	45000	0		28000		{"PPN":0,"Service Charge":0,"PB1":0}		0
450	2020-09-23 23:22:57.862625	SL-20200923-0000006	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0017","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic6","modifier_option":"","modifier_price":0,"number_orders":2,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	130000	0	130000	Uang Pas	Invalid Date	130000	0		110000		{"PPN":0,"Service Charge":0,"PB1":0}		0
451	2020-09-23 23:23:27.339739	SL-20200923-0000007	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0019","name":"KOPI TIRAMISU","category":"Lainnya","variant":"Kopi tiramisu nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":0,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		0		{"PPN":0,"Service Charge":0,"PB1":0}		0
452	2020-09-23 23:25:00.113457	SL-20200923-0000008	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0002","name":"WIRE","category":"Lainnya","variant":"mini alien","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":20000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Uang Pas	Invalid Date	30000	0		20000		{"PPN":0,"Service Charge":0,"PB1":0}		0
464	2020-09-24 11:12:45.831794	SL-20200924-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0016","name":"COTTON","category":"Lainnya","variant":"KENDO","modifier_option":"","modifier_price":0,"number_orders":12,"buy_cost":26000,"sell_cost":40000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	480000	0	480000	BCA	Invalid Date	480000	0		312000		{"PPN":0,"Service Charge":0,"PB1":0}		0
471	2020-09-24 17:23:17.6226	SL-20200924-0000002	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0016","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	BCA	Invalid Date	65000	0		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
474	2020-09-24 20:02:34.874458	SL-20200924-0000003	+6282264291947	OTL-001	Bawa Pulang	CUS-0279	[{"sku_id":"FP-0006","name":"BOMBER","category":"Lainnya","variant":"melon susu","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Uang Pas	Invalid Date	45000	0		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
476	2020-09-24 20:45:46.15806	SL-20200924-0000004	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0016","name":"COTTON","category":"Lainnya","variant":"KENDO","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":26000,"sell_cost":40000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	40000	0	40000	Uang Pas	Invalid Date	40000	0		26000		{"PPN":0,"Service Charge":0,"PB1":0}		0
480	2020-09-24 21:38:41.496111	SL-20200924-0000005	+6282264291947	OTL-001	Bawa Pulang	CUS-0427	[{"sku_id":"FP-0016","name":"COTTON","category":"Lainnya","variant":"KENDO","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":26000,"sell_cost":40000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	40000	0	40000	BCA	Invalid Date	40000	0		26000		{"PPN":0,"Service Charge":0,"PB1":0}		0
524	2020-09-26 15:43:11.847011	SL-20200926-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0243	[{"sku_id":"FP-0019","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":0,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		0		{"PPN":0,"Service Charge":0,"PB1":0}		0
527	2020-09-26 19:33:43.8008	SL-20200926-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0421	[{"sku_id":"FP-0002","name":"WIRE","category":"Lainnya","variant":"mini alien","modifier_option":"","modifier_price":0,"number_orders":2,"buy_cost":20000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	60000	0	60000	Tunai	Invalid Date	60000	40000		40000		{"PPN":0,"Service Charge":0,"PB1":0}		0
572	2020-09-27 15:24:36.845036	SL-20200927-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0011","name":"Mister Moron","category":"Lainnya","variant":"Nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":140000,"sell_cost":160000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	160000	0	160000	Uang Pas	Invalid Date	160000	0		140000		{"PPN":0,"Service Charge":0,"PB1":0}		0
578	2020-09-27 23:46:25.559826	SL-20200927-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0243	[{"sku_id":"FP-0006","name":"BOMBER","category":"Lainnya","variant":"melon susu","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Tunai	Invalid Date	45000	55000		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
579	2020-09-27 23:47:20.347895	SL-20200927-0000003	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0015","name":"CHARGER AWT","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":35000,"sell_cost":50000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	50000	0	50000	Uang Pas	Invalid Date	50000	0		35000		{"PPN":0,"Service Charge":0,"PB1":0}		0
580	2020-09-27 23:48:15.612306	SL-20200927-0000004	+6282264291947	OTL-001	Bawa Pulang	CUS-0382	[{"sku_id":"FP-0010","name":"Mister Moron","category":"Lainnya","variant":"Nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":140000,"sell_cost":160000,"units":"Buah","diskon":0,"tax":"","salestype_up":0,"notes":""}]	160000	0	160000	Uang Pas	Invalid Date	160000	0		140000		{"PPN":0,"Service Charge":0,"PB1":0}		0
656	2020-09-28 16:38:50.737826	SL-20200928-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0294	[{"sku_id":"FP-0014","name":"BANANA LICIOUS","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":107000,"sell_cost":120000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	120000	0	120000	Uang Pas	Invalid Date	120000	0		107000		{"PPN":0,"Service Charge":0,"PB1":0}		0
719	2020-09-28 20:56:14.333221	SL-20200928-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0420	[{"sku_id":"FP-0015","name":"CHARGER AWT","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":35000,"sell_cost":50000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	50000	0	50000	Tunai	Invalid Date	50000	50000		35000		{"PPN":0,"Service Charge":0,"PB1":0}		0
799	2020-09-29 19:24:42.487429	SL-20200929-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0428	[{"sku_id":"FP-0008","name":"BOMBER","category":"Lainnya","variant":"OB Strawberry","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":28000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""},{"sku_id":"FP-0003","name":"WIRE","category":"Lainnya","variant":"fused clapton","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":10000,"sell_cost":20000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Tunai	Invalid Date	65000	35000		38000		{"PPN":0,"Service Charge":0,"PB1":0}		0
802	2020-09-29 19:42:52.216785	SL-20200929-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0429	[{"sku_id":"FP-0005","name":"BOMBER","category":"Lainnya","variant":"Mangga","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Tunai	Invalid Date	45000	5000		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
942	2020-09-30 18:53:48.346678	SL-20200930-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0006","name":"BOMBER","category":"Lainnya","variant":"melon susu","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""},{"sku_id":"FP-0007","name":"BOMBER","category":"Lainnya","variant":"Anggur","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""},{"sku_id":"FP-0015","name":"CHARGER AWT","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":35000,"sell_cost":50000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	140000	0	140000	Uang Pas	Invalid Date	140000	0		85000		{"PPN":0,"Service Charge":0,"PB1":0}		0
947	2020-09-30 20:33:53.872695	SL-20200930-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0430	[{"sku_id":"FP-0001","name":"WIRE","category":"Lainnya","variant":"baby alien","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":15000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Tunai	Invalid Date	30000	20000		15000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1000	2020-10-01 13:51:46.140779	SL-20201001-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0379	[{"sku_id":"FP-0018","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":0,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		0		{"PPN":0,"Service Charge":0,"PB1":0}		0
1064	2020-10-01 19:34:20.223924	SL-20201001-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0431	[{"sku_id":"FP-0018","name":"KOPI TIRAMISU","category":"Lainnya","variant":"Kopi tiramisu nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Tunai	Invalid Date	65000	35000		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1065	2020-10-01 19:39:28.919591	SL-20201001-0000003	+6282264291947	OTL-001	Bawa Pulang	CUS-0285	[{"sku_id":"FP-0022","name":"glass RTA ZEUS X /SULTAN X","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":12000,"sell_cost":30000,"units":"Butir","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Uang Pas	Invalid Date	30000	0		12000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1149	2020-10-02 14:40:00.958282	SL-20201002-0000001	+6282264291947	OTL-001	Bawa Pulang	Belum di Set	[{"sku_id":"FP-0006","name":"BOMBER","category":"Lainnya","variant":"melon susu","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Uang Pas	Invalid Date	45000	0		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1174	2020-10-02 17:15:07.351469	SL-20201002-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0417	[{"sku_id":"FP-0001","name":"WIRE","category":"Lainnya","variant":"baby alien","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":15000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Uang Pas	Invalid Date	30000	0		15000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1189	2020-10-02 21:11:43.350705	SL-20201002-0000003	+6282264291947	OTL-001	Bawa Pulang	CUS-0407	[{"sku_id":"FP-0010","name":"Mister Moron","category":"Lainnya","variant":"Nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":140000,"sell_cost":160000,"units":"Buah","diskon":0,"tax":"","salestype_up":0,"notes":""}]	160000	0	160000	Uang Pas	Invalid Date	160000	0		140000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1190	2020-10-02 21:13:51.947141	SL-20201002-0000004	+6282264291947	OTL-001	Bawa Pulang	CUS-0433	[{"sku_id":"FP-0019","name":"NGEBOEL","category":"Lainnya","variant":"Banabeast nic6","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1231	2020-10-03 10:08:16.921095	SL-20201003-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0294	[{"sku_id":"FP-0017","name":"KENDO VAPE COTTON ","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":26000,"sell_cost":40000,"units":"Pack","diskon":0,"tax":"","salestype_up":0,"notes":""}]	40000	0	40000	Tunai	Invalid Date	40000	10000		26000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1322	2020-10-03 16:54:14.928895	SL-20201003-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0351	[{"sku_id":"FP-0019","name":"NGEBOEL","category":"Lainnya","variant":"Kopi tiramisu nic3","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1325	2020-10-03 19:47:20.869054	SL-20201003-0000003	+6282264291947	OTL-001	Bawa Pulang	CUS-0285	[{"sku_id":"FP-0022","name":"glass RTA ZEUS X /SULTAN X","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":12000,"sell_cost":30000,"units":"Butir","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Tunai	Invalid Date	30000	20000		12000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1394	2020-10-04 13:35:12.617072	SL-20201004-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0294	[{"sku_id":"FP-0021","name":"Fused clapton full TMNi80","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":40000,"units":"Cup","diskon":0,"tax":"","salestype_up":0,"notes":""}]	40000	0	40000	Uang Pas	Invalid Date	40000	0		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1473	2020-10-04 19:56:47.598124	SL-20201004-0000003	+6282264291947	OTL-001	Bawa Pulang	CUS-0351	[{"sku_id":"FP-0002","name":"WIRE","category":"Lainnya","variant":"mini alien","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":20000,"sell_cost":30000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	30000	0	30000	Tunai	Invalid Date	30000	20000		20000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1474	2020-10-04 19:59:52.106442	SL-20201004-0000004	+6282264291947	OTL-001	Bawa Pulang	CUS-0435	[{"sku_id":"FP-0005","name":"BOMBER","category":"Lainnya","variant":"Mangga","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Tunai	Invalid Date	45000	5000		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1497	2020-10-05 08:32:21.835207	SL-20201005-0000001	+6281932809265	OTL-001	Bawa Pulang	CUS-0191	[{"sku_id":"","name":"Kentang","category":"Manual Input","variant":"","modifier_option":"","modifier_price":0,"number_orders":4,"buy_cost":15000,"sell_cost":36000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""},{"sku_id":"","name":"Cimol","category":"Manual Input","variant":"","modifier_option":"","modifier_price":0,"number_orders":3,"buy_cost":7000,"sell_cost":36000,"units":"Pieces","diskon":0,"tax":"","salestype_up":0,"notes":""}]	252000	0	252000	Uang Pas	Invalid Date	252000	0		81000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1433	2020-10-04 16:18:40.308447	SL-20201004-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0028	[{"sku_id":"FP-0025","name":"Pesanan Mbah Pringgo","category":"Jasa","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":180000,"sell_cost":200000,"units":"Batang","diskon":0,"tax":"","salestype_up":0,"notes":""}]	200000	0	200000	Jumlah Hari Lainnya	05-Oct-2020	200000	0	pembayaran buat bayar hutang😁	180000	05-Oct-2020 10:11:07	{"PPN":0,"Service Charge":0,"PB1":0}		0
1579	2020-10-05 20:19:43.368932	SL-20201005-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0243	[{"sku_id":"FP-0008","name":"BOMBER","category":"Lainnya","variant":"OB Strawberry","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":28000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Tunai	Invalid Date	45000	5000		28000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1602	2020-10-06 14:17:43.740488	SL-20201006-0000001	+6282264291947	OTL-001	Bawa Pulang	CUS-0436	[{"sku_id":"FP-0005","name":"BOMBER","category":"Lainnya","variant":"Mangga","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":25000,"sell_cost":45000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	45000	0	45000	Tunai	Invalid Date	45000	5000		25000		{"PPN":0,"Service Charge":0,"PB1":0}		0
1608	2020-10-06 19:22:26.346481	SL-20201006-0000002	+6282264291947	OTL-001	Bawa Pulang	CUS-0317	[{"sku_id":"FP-0022","name":"NGEBOEL Banabeast nic6","category":"Lainnya","variant":"","modifier_option":"","modifier_price":0,"number_orders":1,"buy_cost":55000,"sell_cost":65000,"units":"Botol","diskon":0,"tax":"","salestype_up":0,"notes":""}]	65000	0	65000	Uang Pas	Invalid Date	65000	0		55000		{"PPN":0,"Service Charge":0,"PB1":0}		0
\.


--
-- Data for Name: saved_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.saved_orders (id, create_dtm, user_id, outlet_id, saved_orders_id, name, phone, orders, table_id) FROM stdin;
761	2020-10-06 18:17:00.541655	+6282264291947	OTL-001	#20100618165472	coil tambak	6285314138712	{"tableNumber":"","reservationDate":"","deliveryDate":"","guestNumber":0,"products":[{"modifier_option":"","modifier_price":0,"notes":"","number_orders":1,"basic_sell_cost":260000,"taxInfo":[],"discount_info":{"name":"","amount":0},"id":"3962","create_dtm":"2020-10-06T18:15:42.002976Z","sku_id":"FP-0023","user_id":"+6282264291947","outlet_id":"OTL-001","name":"Pesanan BOS FAUZY","category":"Jasa","variant":"","units":"Boks","weight":0,"quantity":1,"minimum_quantity":0,"description":"RDA MUSE & TM FUSED CLAPTON","buy_cost":245000,"sell_cost":260000,"modifiers_id":"","images":["https://firebasestorage.googleapis.com/v0/b/artaka-mpos.appspot.com/o/300060b3-9a0b-4ceb-8ede-3212dabb786a.jpg?alt=media&token=30a4712a-2add-4bbd-90bd-6687b972e41d"],"rawmaterial":[],"is_stock_tracked":"Yes","number_sold":0,"outlets":["OTL-001"],"buy_cost_discounted":245000,"is_active":"Yes","wholesaler_cost":[],"salestype_up":0}],"customer":{"id":"55001","create_dtm":"2020-09-23T16:33:10.164807Z","customer_id":"CUS-0317","user_id":"+6282264291947","outlet_id":"OTL-001","name":"coil tambak","email":"","phone":"6285314138712","datebirth":"","gender":"","address":"","city":"","province":"-","images":["https://style.anu.edu.au/_anu/4/images/placeholders/person.png"],"points_balance":0},"fromWhatsApp":1,"ongkir":0,"salesType":"WhatsApp","expedition":"Diantar Oleh Penjual","subTotal":260000,"totalBill":260000,"totalTax":0,"totalServiceCharge":0,"totalPB1":0,"totalDiskon":0,"totalWeight":0,"totalBuyCost":245000,"bankAccount":"||"}	
\.


--
-- Data for Name: subscribers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.subscribers (id, create_dtm, user_id, email, owner_name, secret_password, fcm_token, idcard_name, idcard_number, bank_holder_name, bank_name, bank_account, idcard_image, referral_code) FROM stdin;
47	2020-08-06 11:22:16.420747	+62811196196		Iim Rusyamsi	okoce2518	ffDNcEUpSRSnjzABisxEjI:APA91bFKnz8BrvHSqwP5t9PlMZGeffW5rYVJScajBAJBvV7vUv0iPmTTqK0pr33GE6t1oubwowetfAaeIr6sMmYZYUdKXZAi2_nJoGMupyj8sPsa5odYKWH4jzT7QWrPbANyIhILkjnU						["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	
37	2020-07-26 17:24:47.231569	+6281311666268	Mohfajar.173@gmail.con	Moh Fajar	sasha4595							["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	
54	2020-08-14 13:47:19.978995	+628116261101	ulindatobing@gmail.com	ULINDA LUMBANTOBING	gaberaulito1703							["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	
205	2020-09-22 22:25:32.771028	+6282264291947	moemoebodo@gmail.com	MODHENK	291947	c2XyqUlXSNay6Cfssw_upX:APA91bG52oD3LawiBIEtJk5k2X8vNFb56vgYCIyqgULhXvjtRFGI-AC94DoGW7Fd2vA95h6sRzhtUzqdCds8VvGbPIj5aT__YsJ9RabyfBltYEjscgSqjz4mO9EpI_VageaMM36y1PFx			||	||	||	["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	
336	2020-10-05 08:16:00.376981	+6281932809265	ariez.n3@gmail.com	Haris Maulana	123456	dtc7pWWUTQ62qOcucSjfbp:APA91bEQjT1-bXJmgY_nSC5exjaaKpADSSoRczSkTCrAqzyWAFfQWY0EzRnxUgZgqhaE2DZe8kNNxoc8EvyNu3hQCM5QorNa1PE6X1V3ieKdTuTbrasoAzqSOGNMbipuezKHtauTfn6_						["https://www.generationsforpeace.org/wp-content/uploads/2018/07/empty.jpg"]	
\.


--
-- Name: admins_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.admins_id_seq', 1, false);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.posts_id_seq', 1, false);


--
-- Name: admins admins_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_email_key UNIQUE (email);


--
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- Name: admins admins_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_username_key UNIQUE (username);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

