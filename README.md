# Train API

This is a public API meant for be used publicly for frontends.
This is also a good project for backend developers that can only
code CRUD stuff. Now, you have to think in states, as there are
no database to begin with.

I intentionally uses real train stations and coordinates from those
that exists in Jakarta to mimic a real train arrival/departure
behavior. Please do not use this API and advertise it
as if this is a real train schedule. It's definitely not.
BUT, if you are advertising the API as a real train schedule,
please stop. Don't do that.

THIS API IS MEANT FOR EDUCATIONAL PURPOSE.
I AM NOT RESPONSIBLE FOR ANY ABUSE IN THE NAME OF USING THIS PROGRAM.
NOR I DO NOT GAIN ANY MONEY FROM CREATING THIS PROGRAM.

The endpoints are pretty simple, you are most welcome to open a PR
and submit a new endpoint in case you need them:

```http request
### Get all lines
GET http://localhost:5000/lines
Accept: application/json

### Get lines by ID
GET http://localhost:5000/lines?id=1
Accept: application/json

### Get all stations
GET http://localhost:5000/stations
Accept: application/json

### Get station by ID
GET http://localhost:5000/stations?id=1
Accept: application/json

### Get all trains
GET http://localhost:5000/trains
Accept: application/json

### Get trains by ID
GET http://localhost:5000/trains?id=100
Accept: application/json
```

## License

```
   Copyright 2022 Reinaldy Rafli

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

See [LICENSE](./LICENSE)