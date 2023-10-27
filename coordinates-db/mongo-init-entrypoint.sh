mongosh --host mongo:27017 -u root -p root --eval "
	rs.initiate({
		_id: 'rs0',
		members: [
			{_id:0, host: 'mongo'},
		]});
          
        while (rs.status().ok == 0) sleep(1000);
"

sleep 3

mongosh --host mongo:27017 -u root -p root --eval "
	db = db.getSiblingDB('VEHICLE-TRACKING')
	db.createCollection(
		'vehicle-status',
		{
		   timeseries: {
			  timeField: 'ts',
			  metaField: 'meta',
			  granularity: 'seconds'
		   }
		}
	)
"
	
sleep 1

mongosh --host mongo:27017 -u root -p root --eval "
	db = db.getSiblingDB('VEHICLE-TRACKING')
	db.getCollection('vehicle-status').createIndex({'meta.vehicle_id':1, 'ts':-1});
	db.getCollection('vehicle-status').createIndex({'meta.route_id':1});
	db.getCollection('vehicle-status').createIndex({'meta.user_id':1, 'ts':-1});
"