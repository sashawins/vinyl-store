INSERT INTO public.artists (id,name,country,created_at) VALUES
	 ('5bd48594-2d14-466a-a3e1-ea15414389b3'::uuid,'Radiohead','UK','2025-06-17 21:47:26.248245+03'),
	 ('5c54db46-59dc-4912-b61f-cd1475545415'::uuid,'Kendrick Lamar','USA','2025-06-17 21:47:26.248245+03'),
	 ('a08bd006-7b61-450e-a265-d24d394c31d2'::uuid,'Pink Floyd','UK','2025-06-30 02:22:56.608821+03'),
	 ('55433ef6-7dc7-41ef-b09b-89f831d00875'::uuid,'Madvillain','USA','2025-06-30 02:28:04.435874+03'),
	 ('582a1854-a443-4904-a45b-409cf303b0ac'::uuid,'King Crimson','UK','2025-06-30 02:28:23.765058+03'),
	 ('8a4a5c67-d3af-4726-a1f7-0ca1975a7cd8'::uuid,'My Bloody Valentine','Ireland','2025-06-30 02:28:51.266736+03');

INSERT INTO public.genres (id,name,created_at) VALUES
	 ('a71fef37-26a7-4b95-9605-84bfcb745f86'::uuid,'Alternative','2025-06-17 21:47:26.248245+03'),
	 ('60ccb585-4103-4fd5-9f32-57e8462a208c'::uuid,'Hip-Hop','2025-06-17 21:47:26.248245+03'),
	 ('33684b3c-57c1-4357-8536-52c45a272c98'::uuid,'Electronic','2025-06-17 21:47:26.248245+03'),
	 ('6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'Rock','2025-06-17 21:47:26.248245+03'),
	 ('3495be5e-fb03-41b7-b45f-38f963188888'::uuid,'Soul','2025-06-17 21:47:26.248245+03'),
	 ('40010ca0-9759-4a12-be04-54151c255711'::uuid,'Indie Rock','2025-06-17 21:47:26.248245+03'),
	 ('97dff092-01ac-4def-8fb9-2dbb3cfd6110'::uuid,'Folk','2025-06-17 21:47:26.248245+03'),
	 ('7ce25b47-dd3a-4680-9710-3618192f6c9c'::uuid,'Jazz','2025-06-17 21:47:26.248245+03'),
	 ('2c4d57f0-3865-4019-9cf7-d28253ea4735'::uuid,'Pop','2025-06-17 21:47:26.248245+03'),
	 ('97116160-b01a-440e-ad80-ffc304a083c0'::uuid,'Electronic','2025-06-30 02:44:52.497957+03');

INSERT INTO public.news (id,title,"content",published_at) VALUES
	 ('0ea41ffd-b91a-4236-b72d-888efaa80d21'::uuid,'Открытие в Москве','Наш магазин теперь есть и в Москве!<br>адрес магазина: пр. Вернадского, 78, м. Юго-западная.<br>режим работы: пн—вс 15:00 – 22:00<br>телефон: +7 (495) 005-05-05<br>telegram: @accessrecordstore_msc','2025-06-30 03:29:06.869324+03'),
	 ('71a0c77b-d0b9-4180-a406-a558afb81776'::uuid,'Переезд магазина в Екатеринбурге','дорогие друзья, наш магазин на ул. Ленина 13-15 в Екатеринбурге работает до 30 июня включительно, после чего мы будем закрыты и начнем переезд в новое место.<br>до встречи!','2025-06-30 03:32:19.996753+03');

INSERT INTO public.order_items (id,order_id,vinyl_id,quantity,unit_price) VALUES
	 ('b4260333-51c8-46c8-af89-681170670f69'::uuid,'d5584627-ad38-4a5a-bdcd-40aaf4967831'::uuid,'07a9352c-7ff1-414f-8210-fd93a61cf517'::uuid,2,1200.0),
	 ('b81e522f-e751-45d7-8a60-61ef7e29ebf9'::uuid,'d5584627-ad38-4a5a-bdcd-40aaf4967831'::uuid,'1b0b696b-fd3a-49c7-aa85-a6ea174cd4dc'::uuid,1,2990.0),
	 ('5dafbb8a-e4cf-44d7-b018-dc6d553b74c2'::uuid,'16a89733-0995-485f-a5d6-0de902114bec'::uuid,'deb5f449-7ae5-42a0-b28c-fdad2efb1bfa'::uuid,1,1990.0),
	 ('ebb33518-f2c5-456f-8140-c9c17f4e5f68'::uuid,'16a89733-0995-485f-a5d6-0de902114bec'::uuid,'bdaa2245-e43d-4b1e-88cf-122fe3ecf1f6'::uuid,1,5990.0),
	 ('af2ece61-cb01-40f3-99c5-d3572476d5da'::uuid,'16a89733-0995-485f-a5d6-0de902114bec'::uuid,'ba52b279-79c8-4d54-ad21-3709a80148e2'::uuid,1,2450.0);

INSERT INTO public.orders (id,user_id,status,total_amount,created_at) VALUES
	 ('d5584627-ad38-4a5a-bdcd-40aaf4967831'::uuid,'bdaeb881-0d2b-45ff-b9df-a5e021dd9677'::uuid,'pending',5390.0,'2025-06-30 03:44:33.658806+03'),
	 ('16a89733-0995-485f-a5d6-0de902114bec'::uuid,'524da01b-4582-40b2-aa93-d589e82a1fd7'::uuid,'pending',10430.0,'2025-06-30 03:48:18.569782+03');

INSERT INTO public.vinyls (id,title,price,stock_count,cover_url,artist_id,genre_id,created_at) VALUES
	 ('deb5f449-7ae5-42a0-b28c-fdad2efb1bfa'::uuid,'To Pimp a Butterfly',1990.0,8,'https://is1-ssl.mzstatic.com/image/thumb/Music112/v4/0d/ae/61/0dae6140-d4af-d0df-eae0-3c92eb392a33/15UMGIM11922.rgb.jpg/600x600bb.jpg','5c54db46-59dc-4912-b61f-cd1475545415'::uuid,'60ccb585-4103-4fd5-9f32-57e8462a208c'::uuid,'2025-06-17 21:47:26.248245+03'),
	 ('6f541867-f7ea-43de-a123-9d2f9d041713'::uuid,'OK Computer',1200.0,4,'https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/07/60/ba/0760ba0f-148c-b18f-d0ff-169ee96f3af5/634904078164.png/600x600bb.jpg','5bd48594-2d14-466a-a3e1-ea15414389b3'::uuid,'a71fef37-26a7-4b95-9605-84bfcb745f86'::uuid,'2025-06-17 21:47:26.248245+03'),
	 ('1b0b696b-fd3a-49c7-aa85-a6ea174cd4dc'::uuid,'Kid A',2990.0,5,'https://is1-ssl.mzstatic.com/image/thumb/Music122/v4/bd/8e/13/bd8e1358-b367-a689-cb84-cebd0b067dc4/634904078263.png/600x600bb.jpg','59087fa8-2fc8-473c-bf90-13aa6d8e006a'::uuid,'33684b3c-57c1-4357-8536-52c45a272c98'::uuid,'2025-06-17 21:47:26.248245+03'),
	 ('07a9352c-7ff1-414f-8210-fd93a61cf517'::uuid,'In Rainbows',1200.0,5,'https://is1-ssl.mzstatic.com/image/thumb/Music126/v4/dd/50/c7/dd50c790-99ac-d3d0-5ab8-e3891fb8fd52/634904032463.png/600x600bb.jpg','5bd48594-2d14-466a-a3e1-ea15414389b3'::uuid,'6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'2025-06-30 02:31:01.783645+03'),
	 ('025c5a46-262d-4a52-8fd2-051d0654307c'::uuid,'Wish You Were Here',1200.0,7,'https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/49/ab/fe/49abfef6-0cd9-aa1f-05c3-3eb85d3fe3f5/886445635843.jpg/600x600bb.jpg','a08bd006-7b61-450e-a265-d24d394c31d2'::uuid,'6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'2025-06-30 02:33:25.10584+03'),
	 ('3f648027-edde-453e-913e-27244d648e36'::uuid,'The Dark Side of the Moon',2450.0,12,'https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/3c/1b/a9/3c1ba9e1-15b1-03b3-3bfd-09dbd9f1705b/dj.mggvbaou.jpg/600x600bb.jpg','a08bd006-7b61-450e-a265-d24d394c31d2'::uuid,'6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'2025-06-30 02:33:46.337119+03'),
	 ('bdaa2245-e43d-4b1e-88cf-122fe3ecf1f6'::uuid,'Madvillainy',5990.0,3,'https://is1-ssl.mzstatic.com/image/thumb/Music211/v4/2d/33/99/2d33992f-4550-269a-bab2-a2558996e5c3/816.jpg/600x600bb.jpg','55433ef6-7dc7-41ef-b09b-89f831d00875'::uuid,'60ccb585-4103-4fd5-9f32-57e8462a208c'::uuid,'2025-06-30 02:34:27.36883+03'),
	 ('0a412103-b9d7-4a32-92d5-b68c7a2c8bbe'::uuid,'In the Court of the Crimson King',4700.0,7,'https://is1-ssl.mzstatic.com/image/thumb/Music5/v4/2f/c7/19/2fc71988-6871-be2c-6731-a3d0f2a6b232/Court_2500px.jpg/600x600bb.jpg','582a1854-a443-4904-a45b-409cf303b0ac'::uuid,'6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'2025-06-30 02:35:34.005214+03'),
	 ('ba52b279-79c8-4d54-ad21-3709a80148e2'::uuid,'good kid, m.A.A.d city',2450.0,3,'https://is1-ssl.mzstatic.com/image/thumb/Music112/v4/86/d2/5d/86d25d08-2db3-b7f1-d688-96bc78719c28/12UMGIM52989.rgb.jpg/600x600bb.jpg','5c54db46-59dc-4912-b61f-cd1475545415'::uuid,'60ccb585-4103-4fd5-9f32-57e8462a208c'::uuid,'2025-06-30 02:36:18.934848+03'),
	 ('d2b6cbb0-6c65-4b68-838e-26679bf59e7c'::uuid,'Loveless',4700.0,5,'https://is1-ssl.mzstatic.com/image/thumb/Music125/v4/3b/4b/45/3b4b459a-84ed-f03b-0f3c-85fd145a5108/mzi.oewhanoe.jpg/600x600bb.jpg','8a4a5c67-d3af-4726-a1f7-0ca1975a7cd8'::uuid,'6ed53afc-a614-4d5a-a82a-95afd4e2f56f'::uuid,'2025-06-30 02:37:17.954778+03');
