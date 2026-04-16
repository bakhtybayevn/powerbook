-- +goose Up
-- +goose StatementBegin

INSERT INTO users (id, email, display_name, password_hash, streak_current_days, streak_last_date, total_minutes, xp, telegram_handle) VALUES
('0a8f4f51-108b-5c18-8108-1711554bc8f2','alice.turner@example.com','Alice Turner','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',12,'2026-04-15',2400,850,'@alice_reads'),
('3e140cd5-2531-5922-91f5-907f7bb3c2cf','bob.martinez@example.com','Bob Martinez','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',8,'2026-04-15',1800,620,'@bobmartinez'),
('9453a05f-a579-51ae-9a6e-e896c25c5f6b','charlie.kim@example.com','Charlie Kim','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',15,'2026-04-15',3200,1200,'@charlie_books'),
('f801c68b-7720-5109-bd39-32525ab0997e','diana.patel@example.com','Diana Patel','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',5,'2026-04-15',900,310,'@dianareads'),
('30ed9424-ee2d-5b8e-952b-1a8ad2477dac','ethan.jones@example.com','Ethan Jones','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',20,'2026-04-15',4100,1600,'@ethanj'),
('229b7b3e-9b4b-5364-a7fc-b81a54765f11','fiona.wang@example.com','Fiona Wang','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',3,'2026-04-15',600,180,'@fionaw'),
('e41a8000-86c1-5ecd-bde5-2aa293897fd8','george.ali@example.com','George Ali','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',10,'2026-04-15',1500,520,'@georgeali'),
('fefd705f-f26d-5632-a7d0-4d477fa875c3','hannah.lee@example.com','Hannah Lee','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',7,'2026-04-15',1200,430,'@hannahlee'),
('23044b8a-5574-5952-860c-2792dbbdd0b4','isaac.brown@example.com','Isaac Brown','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',25,'2026-04-15',5000,2100,'@isaacb'),
('23d284e9-6cca-563d-99b0-467da67bd6dd','julia.chen@example.com','Julia Chen','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',14,'2026-04-15',2800,950,'@juliachen'),
('4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5','kevin.singh@example.com','Kevin Singh','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',2,'2026-04-15',400,120,'@kevins'),
('b0550926-0d80-5051-b39f-d3b4516128b4','lana.davis@example.com','Lana Davis','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',18,'2026-04-15',3600,1400,'@lanadavis'),
('c5358607-4a2f-5b0d-9ab2-d67c989eb48d','mike.zhao@example.com','Mike Zhao','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',6,'2026-04-15',1000,350,'@mikezhao'),
('9d312b37-d157-5a85-a8f2-77409fb2cd57','nora.wilson@example.com','Nora Wilson','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',11,'2026-04-15',2100,720,'@noraw'),
('7ada812c-321e-5f9c-9f7c-758d0cac347c','oliver.park@example.com','Oliver Park','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',4,'2026-04-15',700,240,'@oliverp'),
('c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3','penny.lopez@example.com','Penny Lopez','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',9,'2026-04-15',1600,560,'@pennylopez'),
('ba598a58-7f74-508f-b721-262b4f28d7a1','quinn.taylor@example.com','Quinn Taylor','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',16,'2026-04-15',3400,1300,'@quinnt'),
('fd43bff3-6e56-5b08-9c6b-25da7d177034','ryan.kumar@example.com','Ryan Kumar','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',1,'2026-04-15',300,80,'@ryankumar'),
('0c055f3f-3a7d-57ad-886e-dc6a53ef0cef','sara.garcia@example.com','Sara Garcia','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',22,'2026-04-15',4500,1800,'@saragarcia'),
('8d951775-b231-504d-b8fb-df6f55db6098','tom.wright@example.com','Tom Wright','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',13,'2026-04-15',2600,880,'@tomwright'),
('bb9479d2-a47c-5d4e-bc42-d35aedde1a4c','uma.pham@example.com','Uma Pham','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',7,'2026-04-15',1100,400,'@umapham'),
('06c3e8b4-7098-52d4-9b6d-af276df38660','victor.ross@example.com','Victor Ross','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',19,'2026-04-15',3800,1500,'@victorr'),
('5543790e-7dfc-5394-a1c0-e67816f856dd','wendy.liu@example.com','Wendy Liu','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',10,'2026-04-15',1900,650,'@wendyliu'),
('3698fb21-2a6f-5a62-aa82-6191b10b0b7e','xander.hill@example.com','Xander Hill','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',3,'2026-04-15',500,160,'@xanderh'),
('6ccb2829-c680-5b43-bc58-d863331ee193','yuki.tanaka@example.com','Yuki Tanaka','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',17,'2026-04-15',3500,1350,'@yukitanaka'),
('5cd242db-31de-5a1c-abac-ca18618e25bf','zara.johnson@example.com','Zara Johnson','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',8,'2026-04-15',1300,470,'@zaraj'),
('3e88bbba-e74e-51ca-858b-59816b934381','adam.smith@example.com','Adam Smith','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',6,'2026-04-15',950,330,'@adamsmith'),
('d419ff55-d1f2-50f2-addf-715fa762ec34','bella.reed@example.com','Bella Reed','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',21,'2026-04-15',4300,1700,'@bellareed'),
('b5bbd54f-66fd-54b5-86e4-0e901e011226','carl.wood@example.com','Carl Wood','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',4,'2026-04-15',650,210,'@carlwood'),
('69fb15ee-411c-5def-8b18-cc5239064705','dina.scott@example.com','Dina Scott','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',11,'2026-04-15',2000,700,'@dinascott'),
('dfbe9a6c-5481-5bbd-8889-d334013205b7','eli.moore@example.com','Eli Moore','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',15,'2026-04-15',3100,1150,'@elimoore'),
('f24c9050-6eb0-5059-8676-5c055048d7fa','faye.clark@example.com','Faye Clark','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',2,'2026-04-15',350,100,'@fayeclark'),
('9b167168-0726-5caf-90d5-3f41c93b2a5b','glen.lewis@example.com','Glen Lewis','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',9,'2026-04-15',1700,590,'@glenlewis'),
('35adf185-4fce-5030-8dff-ac5503e911ff','holly.baker@example.com','Holly Baker','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',13,'2026-04-15',2500,860,'@hollybaker'),
('20a8c79b-e948-5e13-b6ab-c7b2273ac222','ivan.young@example.com','Ivan Young','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',5,'2026-04-15',800,280,'@ivanyoung'),
('840c0259-3f48-591d-9fe0-c2f50502a223','jade.allen@example.com','Jade Allen','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',18,'2026-04-15',3700,1450,'@jadeallen'),
('f4910e79-e8f8-5e87-9565-7094cfc5e9aa','kyle.king@example.com','Kyle King','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',1,'2026-04-15',250,60,'@kyleking'),
('e76d5c51-70af-5a6f-9b29-bb8fa95dd416','luna.green@example.com','Luna Green','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',10,'2026-04-15',1800,630,'@lunagreen'),
('597f1d56-beea-592c-a5db-a34e16ccdc99','max.hall@example.com','Max Hall','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',7,'2026-04-15',1100,390,'@maxhall'),
('7b0724cf-a824-56b2-a011-b1862a958f67','nina.white@example.com','Nina White','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',23,'2026-04-15',4800,2000,'@ninawhite'),
('eef5d933-ef91-5f4c-abf7-4a94b144938e','otto.james@example.com','Otto James','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',6,'2026-04-15',900,310,'@ottojames'),
('9816f97a-4bc0-57a6-a5d9-796d9df62ee6','pia.harris@example.com','Pia Harris','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',14,'2026-04-15',2700,920,'@piaharris'),
('35b72d2b-d4a6-5758-a18e-4a2bd060799d','ravi.nelson@example.com','Ravi Nelson','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',3,'2026-04-15',450,140,'@ravinelson'),
('d3fd51cc-ca0a-5d45-812f-364338b7135a','sue.adams@example.com','Sue Adams','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',12,'2026-04-15',2300,800,'@sueadams'),
('732cb566-fa79-5cde-90b6-f57e1f52010f','tina.brooks@example.com','Tina Brooks','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',8,'2026-04-15',1400,490,'@tinabrooks'),
('c647dfa0-259b-5051-98ad-508b2408c6a3','uri.diaz@example.com','Uri Diaz','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',16,'2026-04-15',3300,1250,'@uridiaz'),
('4d14c1dd-f621-569d-a654-93d6c8a5461a','vera.fox@example.com','Vera Fox','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',4,'2026-04-15',600,200,'@verafox'),
('83de850d-ecb4-5d5d-bda7-436fe79b4f74','will.gray@example.com','Will Gray','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',19,'2026-04-15',3900,1550,'@willgray'),
('4e71dacf-02a9-5650-b1a7-898f7de5c675','xena.hunt@example.com','Xena Hunt','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',5,'2026-04-15',750,260,'@xenahunt'),
('bf75d379-71ad-54ae-9286-702067620fcc','yosef.irwin@example.com','Yosef Irwin','$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',11,'2026-04-15',2050,710,'@yosefirwin')
ON CONFLICT (id) DO NOTHING;

INSERT INTO competitions (id, name, start_date, end_date, status, points_per_minute) VALUES
('b09e43fc-f993-5f86-8acc-38928e63e36e','January Reading Sprint','2026-01-01','2026-01-31','closed',1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO participants (competition_id, user_id, points, days_read, minutes_total, last_log_date) VALUES
('b09e43fc-f993-5f86-8acc-38928e63e36e','0a8f4f51-108b-5c18-8108-1711554bc8f2',2400,28,2400,'2026-01-28'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','3e140cd5-2531-5922-91f5-907f7bb3c2cf',1800,22,1800,'2026-01-29'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','9453a05f-a579-51ae-9a6e-e896c25c5f6b',3200,31,3200,'2026-01-30'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','f801c68b-7720-5109-bd39-32525ab0997e',900,15,900,'2026-01-31'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','30ed9424-ee2d-5b8e-952b-1a8ad2477dac',4100,31,4100,'2026-01-28'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','229b7b3e-9b4b-5364-a7fc-b81a54765f11',600,10,600,'2026-01-29'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','e41a8000-86c1-5ecd-bde5-2aa293897fd8',1500,20,1500,'2026-01-30'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','fefd705f-f26d-5632-a7d0-4d477fa875c3',1200,18,1200,'2026-01-31'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','23044b8a-5574-5952-860c-2792dbbdd0b4',5000,31,5000,'2026-01-28'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','23d284e9-6cca-563d-99b0-467da67bd6dd',2800,25,2800,'2026-01-29'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5',400,8,400,'2026-01-30'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','b0550926-0d80-5051-b39f-d3b4516128b4',3600,30,3600,'2026-01-31'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','c5358607-4a2f-5b0d-9ab2-d67c989eb48d',1000,16,1000,'2026-01-28'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','9d312b37-d157-5a85-a8f2-77409fb2cd57',2100,22,2100,'2026-01-29'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','7ada812c-321e-5f9c-9f7c-758d0cac347c',700,12,700,'2026-01-30'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3',1600,19,1600,'2026-01-31'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','ba598a58-7f74-508f-b721-262b4f28d7a1',3400,29,3400,'2026-01-28'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','fd43bff3-6e56-5b08-9c6b-25da7d177034',300,6,300,'2026-01-29'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef',4500,31,4500,'2026-01-30'),
('b09e43fc-f993-5f86-8acc-38928e63e36e','8d951775-b231-504d-b8fb-df6f55db6098',2600,24,2600,'2026-01-31')
ON CONFLICT (competition_id, user_id) DO NOTHING;

INSERT INTO gift_exchanges (id, competition_id, giver_id, receiver_id, gift_description, giver_confirmed, receiver_confirmed) VALUES
('cd4a4cb3-6a11-5a17-9d90-41fab994b82b','b09e43fc-f993-5f86-8acc-38928e63e36e','23044b8a-5574-5952-860c-2792dbbdd0b4','229b7b3e-9b4b-5364-a7fc-b81a54765f11','The Hitchhiker''s Guide to the Galaxy',true,true),
('7dd5398a-f65a-53ff-b17e-55444cbf64f7','b09e43fc-f993-5f86-8acc-38928e63e36e','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef','fefd705f-f26d-5632-a7d0-4d477fa875c3','A beautiful leather bookmark',true,true),
('3ec176d2-7729-5e2e-a501-b279b0f899a9','b09e43fc-f993-5f86-8acc-38928e63e36e','30ed9424-ee2d-5b8e-952b-1a8ad2477dac','e41a8000-86c1-5ecd-bde5-2aa293897fd8','Kindle gift card $25',true,true),
('a7f3f330-4093-5540-9d7e-1dc56d3eaed3','b09e43fc-f993-5f86-8acc-38928e63e36e','b0550926-0d80-5051-b39f-d3b4516128b4','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5','Dune by Frank Herbert',true,true),
('2c515955-041e-5123-bdc6-eed58723a9ea','b09e43fc-f993-5f86-8acc-38928e63e36e','ba598a58-7f74-508f-b721-262b4f28d7a1','f801c68b-7720-5109-bd39-32525ab0997e','Reading lamp',true,false),
('3c71dae8-c2b3-512b-9332-58b7571cf81e','b09e43fc-f993-5f86-8acc-38928e63e36e','9453a05f-a579-51ae-9a6e-e896c25c5f6b','7ada812c-321e-5f9c-9f7c-758d0cac347c','The Great Gatsby',true,true),
('14e55b51-0039-577b-b820-fd6707874ac2','b09e43fc-f993-5f86-8acc-38928e63e36e','23d284e9-6cca-563d-99b0-467da67bd6dd','fd43bff3-6e56-5b08-9c6b-25da7d177034','Bookstore gift card $20',true,true),
('c9155ba3-07ff-5c02-9621-62e01902f6f0','b09e43fc-f993-5f86-8acc-38928e63e36e','8d951775-b231-504d-b8fb-df6f55db6098','c5358607-4a2f-5b0d-9ab2-d67c989eb48d','Handmade book sleeve',true,true),
('9e63e0b7-0c34-5fa5-99e2-38f3c6685050','b09e43fc-f993-5f86-8acc-38928e63e36e','0a8f4f51-108b-5c18-8108-1711554bc8f2','3e140cd5-2531-5922-91f5-907f7bb3c2cf','Coffee mug with book quote',true,true),
('89a8ddf1-288a-59fe-932e-7ce4de01ea05','b09e43fc-f993-5f86-8acc-38928e63e36e','9d312b37-d157-5a85-a8f2-77409fb2cd57','c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3','Pride and Prejudice',true,true)
ON CONFLICT (competition_id, giver_id) DO NOTHING;

INSERT INTO competitions (id, name, start_date, end_date, status, points_per_minute) VALUES
('3eda2d1b-ad7d-5709-987a-36286157f437','February Book Marathon','2026-02-01','2026-02-28','closed',2)
ON CONFLICT (id) DO NOTHING;

INSERT INTO participants (competition_id, user_id, points, days_read, minutes_total, last_log_date) VALUES
('3eda2d1b-ad7d-5709-987a-36286157f437','0a8f4f51-108b-5c18-8108-1711554bc8f2',4200,26,2100,'2026-02-20'),
('3eda2d1b-ad7d-5709-987a-36286157f437','3e140cd5-2531-5922-91f5-907f7bb3c2cf',3100,20,1550,'2026-02-21'),
('3eda2d1b-ad7d-5709-987a-36286157f437','9453a05f-a579-51ae-9a6e-e896c25c5f6b',5800,28,2900,'2026-02-22'),
('3eda2d1b-ad7d-5709-987a-36286157f437','f801c68b-7720-5109-bd39-32525ab0997e',1600,12,800,'2026-02-23'),
('3eda2d1b-ad7d-5709-987a-36286157f437','30ed9424-ee2d-5b8e-952b-1a8ad2477dac',7200,28,3600,'2026-02-24'),
('3eda2d1b-ad7d-5709-987a-36286157f437','229b7b3e-9b4b-5364-a7fc-b81a54765f11',1000,8,500,'2026-02-25'),
('3eda2d1b-ad7d-5709-987a-36286157f437','e41a8000-86c1-5ecd-bde5-2aa293897fd8',2600,18,1300,'2026-02-26'),
('3eda2d1b-ad7d-5709-987a-36286157f437','fefd705f-f26d-5632-a7d0-4d477fa875c3',2200,16,1100,'2026-02-27'),
('3eda2d1b-ad7d-5709-987a-36286157f437','23044b8a-5574-5952-860c-2792dbbdd0b4',8800,28,4400,'2026-02-28'),
('3eda2d1b-ad7d-5709-987a-36286157f437','23d284e9-6cca-563d-99b0-467da67bd6dd',4800,24,2400,'2026-02-20'),
('3eda2d1b-ad7d-5709-987a-36286157f437','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5',800,6,400,'2026-02-21'),
('3eda2d1b-ad7d-5709-987a-36286157f437','b0550926-0d80-5051-b39f-d3b4516128b4',6400,27,3200,'2026-02-22'),
('3eda2d1b-ad7d-5709-987a-36286157f437','c5358607-4a2f-5b0d-9ab2-d67c989eb48d',1800,14,900,'2026-02-23'),
('3eda2d1b-ad7d-5709-987a-36286157f437','9d312b37-d157-5a85-a8f2-77409fb2cd57',3800,22,1900,'2026-02-24'),
('3eda2d1b-ad7d-5709-987a-36286157f437','7ada812c-321e-5f9c-9f7c-758d0cac347c',1200,10,600,'2026-02-25'),
('3eda2d1b-ad7d-5709-987a-36286157f437','c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3',2800,17,1400,'2026-02-26'),
('3eda2d1b-ad7d-5709-987a-36286157f437','ba598a58-7f74-508f-b721-262b4f28d7a1',6000,27,3000,'2026-02-27'),
('3eda2d1b-ad7d-5709-987a-36286157f437','fd43bff3-6e56-5b08-9c6b-25da7d177034',500,4,250,'2026-02-28'),
('3eda2d1b-ad7d-5709-987a-36286157f437','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef',7800,28,3900,'2026-02-20'),
('3eda2d1b-ad7d-5709-987a-36286157f437','8d951775-b231-504d-b8fb-df6f55db6098',4400,23,2200,'2026-02-21'),
('3eda2d1b-ad7d-5709-987a-36286157f437','bb9479d2-a47c-5d4e-bc42-d35aedde1a4c',2000,15,1000,'2026-02-22'),
('3eda2d1b-ad7d-5709-987a-36286157f437','06c3e8b4-7098-52d4-9b6d-af276df38660',6800,28,3400,'2026-02-23'),
('3eda2d1b-ad7d-5709-987a-36286157f437','5543790e-7dfc-5394-a1c0-e67816f856dd',3400,20,1700,'2026-02-24'),
('3eda2d1b-ad7d-5709-987a-36286157f437','3698fb21-2a6f-5a62-aa82-6191b10b0b7e',900,7,450,'2026-02-25'),
('3eda2d1b-ad7d-5709-987a-36286157f437','6ccb2829-c680-5b43-bc58-d863331ee193',6200,27,3100,'2026-02-26'),
('3eda2d1b-ad7d-5709-987a-36286157f437','5cd242db-31de-5a1c-abac-ca18618e25bf',2400,16,1200,'2026-02-27'),
('3eda2d1b-ad7d-5709-987a-36286157f437','3e88bbba-e74e-51ca-858b-59816b934381',1400,11,700,'2026-02-28'),
('3eda2d1b-ad7d-5709-987a-36286157f437','d419ff55-d1f2-50f2-addf-715fa762ec34',7600,28,3800,'2026-02-20'),
('3eda2d1b-ad7d-5709-987a-36286157f437','b5bbd54f-66fd-54b5-86e4-0e901e011226',1100,9,550,'2026-02-21'),
('3eda2d1b-ad7d-5709-987a-36286157f437','69fb15ee-411c-5def-8b18-cc5239064705',3600,21,1800,'2026-02-22')
ON CONFLICT (competition_id, user_id) DO NOTHING;

INSERT INTO gift_exchanges (id, competition_id, giver_id, receiver_id, gift_description, giver_confirmed, receiver_confirmed) VALUES
('13b47e43-63b9-5954-9fa9-9377a8aae5fe','3eda2d1b-ad7d-5709-987a-36286157f437','23044b8a-5574-5952-860c-2792dbbdd0b4','e41a8000-86c1-5ecd-bde5-2aa293897fd8','1984 by George Orwell',true,true),
('9e1b75de-fb01-5456-a583-d8b27a52c8b0','3eda2d1b-ad7d-5709-987a-36286157f437','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef','bb9479d2-a47c-5d4e-bc42-d35aedde1a4c','Book-themed socks',true,true),
('42224186-353a-51eb-b2f0-2d69c68bf353','3eda2d1b-ad7d-5709-987a-36286157f437','d419ff55-d1f2-50f2-addf-715fa762ec34','f801c68b-7720-5109-bd39-32525ab0997e','Amazon gift card $30',true,true),
('ca1bc7e7-c2ad-5b2e-8cc4-dbac87e3a147','3eda2d1b-ad7d-5709-987a-36286157f437','30ed9424-ee2d-5b8e-952b-1a8ad2477dac','3e88bbba-e74e-51ca-858b-59816b934381','Sapiens by Yuval Noah Harari',true,false),
('b2cfef2b-f72c-5945-a5ca-632e38f3fa6c','3eda2d1b-ad7d-5709-987a-36286157f437','06c3e8b4-7098-52d4-9b6d-af276df38660','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5','Cozy reading blanket',true,true)
ON CONFLICT (competition_id, giver_id) DO NOTHING;

INSERT INTO competitions (id, name, start_date, end_date, status, points_per_minute) VALUES
('bce050a0-b262-503e-a57e-01d3060c9a57','March Madness Reading','2026-03-01','2026-03-31','closed',1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO participants (competition_id, user_id, points, days_read, minutes_total, last_log_date) VALUES
('bce050a0-b262-503e-a57e-01d3060c9a57','0a8f4f51-108b-5c18-8108-1711554bc8f2',2800,26,2800,'2026-03-18'),
('bce050a0-b262-503e-a57e-01d3060c9a57','3e140cd5-2531-5922-91f5-907f7bb3c2cf',2000,20,2000,'2026-03-19'),
('bce050a0-b262-503e-a57e-01d3060c9a57','9453a05f-a579-51ae-9a6e-e896c25c5f6b',3600,30,3600,'2026-03-20'),
('bce050a0-b262-503e-a57e-01d3060c9a57','f801c68b-7720-5109-bd39-32525ab0997e',1100,14,1100,'2026-03-21'),
('bce050a0-b262-503e-a57e-01d3060c9a57','30ed9424-ee2d-5b8e-952b-1a8ad2477dac',4600,31,4600,'2026-03-22'),
('bce050a0-b262-503e-a57e-01d3060c9a57','229b7b3e-9b4b-5364-a7fc-b81a54765f11',700,10,700,'2026-03-23'),
('bce050a0-b262-503e-a57e-01d3060c9a57','e41a8000-86c1-5ecd-bde5-2aa293897fd8',1700,19,1700,'2026-03-24'),
('bce050a0-b262-503e-a57e-01d3060c9a57','fefd705f-f26d-5632-a7d0-4d477fa875c3',1400,17,1400,'2026-03-25'),
('bce050a0-b262-503e-a57e-01d3060c9a57','23044b8a-5574-5952-860c-2792dbbdd0b4',5500,31,5500,'2026-03-26'),
('bce050a0-b262-503e-a57e-01d3060c9a57','23d284e9-6cca-563d-99b0-467da67bd6dd',3200,25,3200,'2026-03-27'),
('bce050a0-b262-503e-a57e-01d3060c9a57','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5',500,8,500,'2026-03-28'),
('bce050a0-b262-503e-a57e-01d3060c9a57','b0550926-0d80-5051-b39f-d3b4516128b4',4000,29,4000,'2026-03-29'),
('bce050a0-b262-503e-a57e-01d3060c9a57','c5358607-4a2f-5b0d-9ab2-d67c989eb48d',1200,16,1200,'2026-03-30'),
('bce050a0-b262-503e-a57e-01d3060c9a57','9d312b37-d157-5a85-a8f2-77409fb2cd57',2400,22,2400,'2026-03-31'),
('bce050a0-b262-503e-a57e-01d3060c9a57','7ada812c-321e-5f9c-9f7c-758d0cac347c',850,12,850,'2026-03-18'),
('bce050a0-b262-503e-a57e-01d3060c9a57','c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3',1800,19,1800,'2026-03-19'),
('bce050a0-b262-503e-a57e-01d3060c9a57','ba598a58-7f74-508f-b721-262b4f28d7a1',3800,29,3800,'2026-03-20'),
('bce050a0-b262-503e-a57e-01d3060c9a57','fd43bff3-6e56-5b08-9c6b-25da7d177034',350,6,350,'2026-03-21'),
('bce050a0-b262-503e-a57e-01d3060c9a57','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef',5000,31,5000,'2026-03-22'),
('bce050a0-b262-503e-a57e-01d3060c9a57','8d951775-b231-504d-b8fb-df6f55db6098',2900,24,2900,'2026-03-23'),
('bce050a0-b262-503e-a57e-01d3060c9a57','bb9479d2-a47c-5d4e-bc42-d35aedde1a4c',1300,15,1300,'2026-03-24'),
('bce050a0-b262-503e-a57e-01d3060c9a57','06c3e8b4-7098-52d4-9b6d-af276df38660',4200,30,4200,'2026-03-25'),
('bce050a0-b262-503e-a57e-01d3060c9a57','5543790e-7dfc-5394-a1c0-e67816f856dd',2100,20,2100,'2026-03-26'),
('bce050a0-b262-503e-a57e-01d3060c9a57','3698fb21-2a6f-5a62-aa82-6191b10b0b7e',600,9,600,'2026-03-27'),
('bce050a0-b262-503e-a57e-01d3060c9a57','6ccb2829-c680-5b43-bc58-d863331ee193',3900,29,3900,'2026-03-28'),
('bce050a0-b262-503e-a57e-01d3060c9a57','5cd242db-31de-5a1c-abac-ca18618e25bf',1600,16,1600,'2026-03-29'),
('bce050a0-b262-503e-a57e-01d3060c9a57','3e88bbba-e74e-51ca-858b-59816b934381',1000,13,1000,'2026-03-30'),
('bce050a0-b262-503e-a57e-01d3060c9a57','d419ff55-d1f2-50f2-addf-715fa762ec34',4400,31,4400,'2026-03-31'),
('bce050a0-b262-503e-a57e-01d3060c9a57','b5bbd54f-66fd-54b5-86e4-0e901e011226',750,10,750,'2026-03-18'),
('bce050a0-b262-503e-a57e-01d3060c9a57','69fb15ee-411c-5def-8b18-cc5239064705',2200,21,2200,'2026-03-19'),
('bce050a0-b262-503e-a57e-01d3060c9a57','dfbe9a6c-5481-5bbd-8889-d334013205b7',3400,28,3400,'2026-03-20'),
('bce050a0-b262-503e-a57e-01d3060c9a57','f24c9050-6eb0-5059-8676-5c055048d7fa',400,7,400,'2026-03-21'),
('bce050a0-b262-503e-a57e-01d3060c9a57','9b167168-0726-5caf-90d5-3f41c93b2a5b',1900,19,1900,'2026-03-22'),
('bce050a0-b262-503e-a57e-01d3060c9a57','35adf185-4fce-5030-8dff-ac5503e911ff',2700,24,2700,'2026-03-23'),
('bce050a0-b262-503e-a57e-01d3060c9a57','20a8c79b-e948-5e13-b6ab-c7b2273ac222',900,12,900,'2026-03-24'),
('bce050a0-b262-503e-a57e-01d3060c9a57','840c0259-3f48-591d-9fe0-c2f50502a223',4100,30,4100,'2026-03-25'),
('bce050a0-b262-503e-a57e-01d3060c9a57','f4910e79-e8f8-5e87-9565-7094cfc5e9aa',300,5,300,'2026-03-26'),
('bce050a0-b262-503e-a57e-01d3060c9a57','e76d5c51-70af-5a6f-9b29-bb8fa95dd416',2000,19,2000,'2026-03-27'),
('bce050a0-b262-503e-a57e-01d3060c9a57','597f1d56-beea-592c-a5db-a34e16ccdc99',1300,15,1300,'2026-03-28'),
('bce050a0-b262-503e-a57e-01d3060c9a57','7b0724cf-a824-56b2-a011-b1862a958f67',5200,31,5200,'2026-03-29')
ON CONFLICT (competition_id, user_id) DO NOTHING;

INSERT INTO gift_exchanges (id, competition_id, giver_id, receiver_id, gift_description, giver_confirmed, receiver_confirmed) VALUES
('7f20d718-5573-5ec1-bc4f-26dd5715c05c','bce050a0-b262-503e-a57e-01d3060c9a57','23044b8a-5574-5952-860c-2792dbbdd0b4','20a8c79b-e948-5e13-b6ab-c7b2273ac222','To Kill a Mockingbird',true,true),
('dc895836-8ec5-5989-968e-c93f6ec0bb10','bce050a0-b262-503e-a57e-01d3060c9a57','7b0724cf-a824-56b2-a011-b1862a958f67','f4910e79-e8f8-5e87-9565-7094cfc5e9aa','Reading journal',true,true),
('a17779e1-31bd-51bc-b0d4-6e7a4d6155c8','bce050a0-b262-503e-a57e-01d3060c9a57','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef','597f1d56-beea-592c-a5db-a34e16ccdc99','Bookstore gift card $15',true,true),
('ca2f4739-7ee8-5f38-bee2-6451cdb2768e','bce050a0-b262-503e-a57e-01d3060c9a57','30ed9424-ee2d-5b8e-952b-1a8ad2477dac','3698fb21-2a6f-5a62-aa82-6191b10b0b7e','The Alchemist',true,true),
('3b98100b-4e7a-5c99-a502-a4385f9b9c68','bce050a0-b262-503e-a57e-01d3060c9a57','d419ff55-d1f2-50f2-addf-715fa762ec34','3e88bbba-e74e-51ca-858b-59816b934381','Tea sampler set',true,false),
('4d2e9701-ebd8-5a4a-9f47-15294be79119','bce050a0-b262-503e-a57e-01d3060c9a57','06c3e8b4-7098-52d4-9b6d-af276df38660','7ada812c-321e-5f9c-9f7c-758d0cac347c','Book light clip-on',true,true),
('2ba164be-8cd4-5041-b9a5-e4d9ac7b57e2','bce050a0-b262-503e-a57e-01d3060c9a57','840c0259-3f48-591d-9fe0-c2f50502a223','f24c9050-6eb0-5059-8676-5c055048d7fa','Handmade bookmark set',true,true),
('b30f8b59-e311-50af-b26a-08078149172c','bce050a0-b262-503e-a57e-01d3060c9a57','b0550926-0d80-5051-b39f-d3b4516128b4','bb9479d2-a47c-5d4e-bc42-d35aedde1a4c','Little Prince illustrated',true,true),
('1bdc051c-a428-5a48-a59a-1017fb37bb76','bce050a0-b262-503e-a57e-01d3060c9a57','6ccb2829-c680-5b43-bc58-d863331ee193','f801c68b-7720-5109-bd39-32525ab0997e','Motivational poster',true,true),
('445809c0-6e8c-54b6-8f67-6b6b3a197a9f','bce050a0-b262-503e-a57e-01d3060c9a57','ba598a58-7f74-508f-b721-262b4f28d7a1','9b167168-0726-5caf-90d5-3f41c93b2a5b','Harry Potter box set',true,true)
ON CONFLICT (competition_id, giver_id) DO NOTHING;

-- +goose StatementEnd

-- +goose Down
DELETE FROM gift_exchanges WHERE competition_id IN ('b09e43fc-f993-5f86-8acc-38928e63e36e','3eda2d1b-ad7d-5709-987a-36286157f437','bce050a0-b262-503e-a57e-01d3060c9a57');
DELETE FROM participants WHERE competition_id IN ('b09e43fc-f993-5f86-8acc-38928e63e36e','3eda2d1b-ad7d-5709-987a-36286157f437','bce050a0-b262-503e-a57e-01d3060c9a57');
DELETE FROM competitions WHERE id IN ('b09e43fc-f993-5f86-8acc-38928e63e36e','3eda2d1b-ad7d-5709-987a-36286157f437','bce050a0-b262-503e-a57e-01d3060c9a57');
DELETE FROM users WHERE id IN ('0a8f4f51-108b-5c18-8108-1711554bc8f2','3e140cd5-2531-5922-91f5-907f7bb3c2cf','9453a05f-a579-51ae-9a6e-e896c25c5f6b','f801c68b-7720-5109-bd39-32525ab0997e','30ed9424-ee2d-5b8e-952b-1a8ad2477dac','229b7b3e-9b4b-5364-a7fc-b81a54765f11','e41a8000-86c1-5ecd-bde5-2aa293897fd8','fefd705f-f26d-5632-a7d0-4d477fa875c3','23044b8a-5574-5952-860c-2792dbbdd0b4','23d284e9-6cca-563d-99b0-467da67bd6dd','4ea713d4-455d-5b8b-93f9-cb66c9a0f7a5','b0550926-0d80-5051-b39f-d3b4516128b4','c5358607-4a2f-5b0d-9ab2-d67c989eb48d','9d312b37-d157-5a85-a8f2-77409fb2cd57','7ada812c-321e-5f9c-9f7c-758d0cac347c','c91214e7-b60a-5cbe-9bb3-3a2ebc2c11a3','ba598a58-7f74-508f-b721-262b4f28d7a1','fd43bff3-6e56-5b08-9c6b-25da7d177034','0c055f3f-3a7d-57ad-886e-dc6a53ef0cef','8d951775-b231-504d-b8fb-df6f55db6098','bb9479d2-a47c-5d4e-bc42-d35aedde1a4c','06c3e8b4-7098-52d4-9b6d-af276df38660','5543790e-7dfc-5394-a1c0-e67816f856dd','3698fb21-2a6f-5a62-aa82-6191b10b0b7e','6ccb2829-c680-5b43-bc58-d863331ee193','5cd242db-31de-5a1c-abac-ca18618e25bf','3e88bbba-e74e-51ca-858b-59816b934381','d419ff55-d1f2-50f2-addf-715fa762ec34','b5bbd54f-66fd-54b5-86e4-0e901e011226','69fb15ee-411c-5def-8b18-cc5239064705','dfbe9a6c-5481-5bbd-8889-d334013205b7','f24c9050-6eb0-5059-8676-5c055048d7fa','9b167168-0726-5caf-90d5-3f41c93b2a5b','35adf185-4fce-5030-8dff-ac5503e911ff','20a8c79b-e948-5e13-b6ab-c7b2273ac222','840c0259-3f48-591d-9fe0-c2f50502a223','f4910e79-e8f8-5e87-9565-7094cfc5e9aa','e76d5c51-70af-5a6f-9b29-bb8fa95dd416','597f1d56-beea-592c-a5db-a34e16ccdc99','7b0724cf-a824-56b2-a011-b1862a958f67','eef5d933-ef91-5f4c-abf7-4a94b144938e','9816f97a-4bc0-57a6-a5d9-796d9df62ee6','35b72d2b-d4a6-5758-a18e-4a2bd060799d','d3fd51cc-ca0a-5d45-812f-364338b7135a','732cb566-fa79-5cde-90b6-f57e1f52010f','c647dfa0-259b-5051-98ad-508b2408c6a3','4d14c1dd-f621-569d-a654-93d6c8a5461a','83de850d-ecb4-5d5d-bda7-436fe79b4f74','4e71dacf-02a9-5650-b1a7-898f7de5c675','bf75d379-71ad-54ae-9286-702067620fcc');
