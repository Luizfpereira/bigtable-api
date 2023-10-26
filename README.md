# bigtable-api

Bigtable API allows users to query data from Bigtable using flexible parameters as datatypes, IDs, dates, regular expressions, count, versions and more.

## Functionalities

- Query data from climate-data table using prefixes (incomplete keys)
- Specify the datatype (weather or forecast)
- Query one area or more for a specific date
- Query one area or more for a range of dates
- Filter your query with parameters:
  - Version
  - Regexp
  - Count

## Environment variables

- PROJECT_ID
- _INSTANCE_ID

## Dependencies

It is necessary to have the Google Application Credentials with the correct permissions to allow the creation of a bigtable client instance.

## Deploy information

This API uses Cloud Build for building in the GCP Platform. When creating the trigger, it is important to inform the Bigtable's _INSTANCE_ID according to the project where the build will be created.

## Installation

clone this repository:
```shell
git clone https://github.com/ZeusAgrotech/bigtable-api.git
```

## Running Transactions App
1. Navigate to the application repository:
```shell
cd bigtable-api
```

2. Inside the repository in root, execute the command:
```shell
go run main.go
```

## Usage

### Routes

Method | Endpoint | Description
------ |--------- | -----------
GET    | /        | check if application is running
GET   | /read/climate-data | Query data from the climate-data table

### Parameters

Parameters can be included in an API request by modifying the URL. This will specify the criteria to determine which records will be returned.

Available parameters include:

- type
- area_id
- date
- version
- regexp
- count

The keys stored in the climate-data table follow the sequece: datatype + Area ID + date. The datatype can be `w (weather)` or `f (forecast)`, the Area ID are composed by the ID of the area, coming with an 'A' as prefix. The date has the format: `YYYY-MM-DD hh:mm:ss`.

Some examples of keys are:

- `w/A327734/2023-10-20 01:00:00`
- `f/A327735/2023-10-25 10:30:00`

The Bigtable does not overwrite data. It means that, everytime a data with a key that is alreay stored in the table is inserted, a new cell is created for that key with a certain timestamp. By default, the API will return only the latest data stored for the keys requested, however, it is possible to query more than one cell for every key by applying the parameter `version`.

Example:

`http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-25 10:30:00&version=3`

## Example Usage

### GET /read/climate-data?type=w

When requesting data from the climate-data table, it is necessary to inform at least a type, which can be: `w (weather)` or `f (forecast)`.

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w'
```

Even though this request queries all weather data in the database, it is not recommended, as it would demand a relevant amount of time to return a result and the response size would be large enough to return some error to the user.

### Read incomplete key (prefix)

After informing the datatype, it is possible to query all keys starting with a certain prefix. For example, the area_id 'A3277' does not exist, so the result will inform every key starting with 'w/A32773'.

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A3277'
```

Response:

Status code: 200 OK
```json
{
  "result": [
        (...)
        {
        "key": "w/A327732/2023-10-18 16:00:00",
        "created": "2023-10-20T03:02:48.844Z",
        "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":34.15,\"temperatureMin\":31.62,\"temperatureMax\":36.08,\"humidityInst\":32.92,\"humidityMin\":27.04,\"humidityMax\":44.88,\"atmosphericPressureInst\":870.59,\"atmosphericPressureMin\":453.5,\"atmosphericPressureMax\":1147.7,\"solarIrradianceInst\":967.94,\"solarIrradianceMin\":164.22,\"solarIrradianceMax\":1068.09,\"solarIrradiation\":952.11,\"rain\":0,\"windSpeedInst\":11.37,\"windDirectionInst\":9.04,\"windSpeedGust\":21.84,\"windDirectionGust\":28.74}}"
        },
        {
        "key": "w/A327732/2023-10-18 16:09:46",
        "created": "2023-10-19T03:02:53.679Z",
        "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":886.23,\"temperatureMin\":453.5,\"temperatureMax\":1147.7,\"humidityInst\":886.23,\"humidityMin\":453.5,\"humidityMax\":1147.7,\"atmosphericPressureInst\":886.23,\"atmosphericPressureMin\":453.5,\"atmosphericPressureMax\":1147.7,\"solarIrradianceInst\":886.23,\"solarIrradianceMin\":453.5,\"solarIrradianceMax\":1147.7,\"solarIrradiation\":922.82,\"rain\":922.82,\"windSpeedInst\":11.37,\"windDirectionInst\":9.04,\"windSpeedGust\":21.84,\"windDirectionGust\":28.74}}"
        },
        {
        "key": "w/A327735/2023-03-16 23:00:00",
        "created": "2023-10-16T20:03:58.315Z",
        "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":21.96,\"temperatureMin\":20.88,\"temperatureMax\":23.2,\"humidityInst\":97.12,\"humidityMin\":91.03,\"humidityMax\":100,\"atmosphericPressureInst\":1128.73,\"atmosphericPressureMin\":592.4,\"atmosphericPressureMax\":1223.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":3.13,\"windDirectionInst\":151.49,\"windSpeedGust\":7.24,\"windDirectionGust\":160.04}}"
        },
        {
        "key": "w/A327735/2023-03-17 00:00:00",
        "created": "2023-10-16T20:03:58.327Z",
        "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":21.88,\"temperatureMin\":20.82,\"temperatureMax\":22.89,\"humidityInst\":96.5,\"humidityMin\":91.8,\"humidityMax\":100,\"atmosphericPressureInst\":1128.78,\"atmosphericPressureMin\":592.4,\"atmosphericPressureMax\":1223.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":1.07,\"windDirectionInst\":195.02,\"windSpeedGust\":9.05,\"windDirectionGust\":200.03}}"
        },
        (...)
  ],
  "status": "success"
}
```

Another example would be to query the datatype 'w', a complete area_id 'A327734' and an incomplete date '2023-10-20'.

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        (...)
        {
            "key": "w/A327734/2023-10-20 20:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.34,\"temperatureMin\":28.58,\"temperatureMax\":37.65,\"humidityInst\":61.44,\"humidityMin\":32.26,\"humidityMax\":72.34,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":123.22,\"solarIrradianceMin\":35.72,\"solarIrradianceMax\":590.85,\"solarIrradiation\":180.42,\"rain\":0,\"windSpeedInst\":21.99,\"windDirectionInst\":154.3,\"windSpeedGust\":45.26,\"windDirectionGust\":199.16}}"
        },
        {
            "key": "w/A327734/2023-10-20 20:09:27",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.34,\"temperatureMin\":28.58,\"temperatureMax\":35.71,\"humidityInst\":30.34,\"humidityMin\":28.58,\"humidityMax\":35.71,\"atmosphericPressureInst\":30.34,\"atmosphericPressureMin\":28.58,\"atmosphericPressureMax\":35.71,\"solarIrradianceInst\":30.34,\"solarIrradianceMin\":28.58,\"solarIrradianceMax\":35.71,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":21.99,\"windDirectionInst\":154.3,\"windSpeedGust\":45.26,\"windDirectionGust\":199.16}}"
        },
        {
            "key": "w/A327734/2023-10-20 21:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":29.05,\"temperatureMin\":27.25,\"temperatureMax\":32.34,\"humidityInst\":64.12,\"humidityMin\":51.29,\"humidityMax\":75.4,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":16.6,\"solarIrradianceMin\":7.58,\"solarIrradianceMax\":259.27,\"solarIrradiation\":61.8,\"rain\":0,\"windSpeedInst\":10.65,\"windDirectionInst\":49.19,\"windSpeedGust\":36.21,\"windDirectionGust\":165.14}}"
        },
        {
            "key": "w/A327734/2023-10-20 21:09:48",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":16.26,\"temperatureMin\":12.04,\"temperatureMax\":12.04,\"humidityInst\":16.26,\"humidityMin\":12.04,\"humidityMax\":12.04,\"atmosphericPressureInst\":16.26,\"atmosphericPressureMin\":12.04,\"atmosphericPressureMax\":12.04,\"solarIrradianceInst\":16.26,\"solarIrradianceMin\":12.04,\"solarIrradianceMax\":12.04,\"solarIrradiation\":67.08,\"rain\":67.08,\"windSpeedInst\":10.65,\"windDirectionInst\":49.19,\"windSpeedGust\":36.21,\"windDirectionGust\":165.14}}"
        },
        {
            "key": "w/A327734/2023-10-20 22:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":27.97,\"temperatureMin\":26.07,\"temperatureMax\":30.02,\"humidityInst\":67.35,\"humidityMin\":57.95,\"humidityMax\":80.42,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":32.98,\"solarIrradiation\":10.3,\"rain\":0,\"windSpeedInst\":2.98,\"windDirectionInst\":15.67,\"windSpeedGust\":23.53,\"windDirectionGust\":35.24}}"
        },
        {
            "key": "w/A327734/2023-10-20 22:10:45",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":0,\"temperatureMin\":0,\"temperatureMax\":0,\"humidityInst\":0,\"humidityMin\":0,\"humidityMax\":0,\"atmosphericPressureInst\":0,\"atmosphericPressureMin\":0,\"atmosphericPressureMax\":0,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":11.21,\"rain\":11.21,\"windSpeedInst\":2.98,\"windDirectionInst\":15.67,\"windSpeedGust\":23.53,\"windDirectionGust\":35.24}}"
        },
    ],
    "status": "success"
}
```

### Read complete key

One can search for a specific area ID and date.

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20 00:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        }
    ],
    "status": "success"
}
```

### Reading more than one area or date

It is possible to do a flexible search, querying one or more areas and one or a range of dates.

When requesting a range of dates, the result will present all dates between the first (inclusive) and the second (exlusive) dates specified.

Example: Query two areas with one date
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-20 00:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327735/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.855Z",
            "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":29.62,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":46.1,\"humidityMin\":29.57,\"humidityMax\":79.32,\"atmosphericPressureInst\":932.73,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        }
    ],
    "status": "success"
}
```

Example: Query one area with a range of dates
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-20 22:00:00,2023-10-21 05:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 23:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":26.57,\"temperatureMin\":24.81,\"temperatureMax\":28.86,\"humidityInst\":73.77,\"humidityMin\":60.93,\"humidityMax\":84.52,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":1.97,\"windDirectionInst\":78.08,\"windSpeedGust\":10.86,\"windDirectionGust\":46.66}}"
        },
        {
            "key": "w/A327734/2023-10-20 23:10:16",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":583.41,\"temperatureMin\":583.4,\"temperatureMax\":1265.6,\"humidityInst\":583.41,\"humidityMin\":583.4,\"humidityMax\":1265.6,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":583.41,\"solarIrradianceMin\":583.4,\"solarIrradianceMax\":1265.6,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":1.97,\"windDirectionInst\":78.08,\"windSpeedGust\":10.86,\"windDirectionGust\":46.66}}"
        },
        {
            "key": "w/A327734/2023-10-21 00:00:00",
            "created": "2023-10-23T03:03:00.022Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":25.94,\"temperatureMin\":24.18,\"temperatureMax\":27.93,\"humidityInst\":78.1,\"humidityMin\":63.89,\"humidityMax\":86.17,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":2.52,\"windDirectionInst\":143.01,\"windSpeedGust\":5.43,\"windDirectionGust\":209.53}}"
        },
        {
            "key": "w/A327734/2023-10-21 00:10:53",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":25.94,\"temperatureMin\":24.73,\"temperatureMax\":27.9,\"humidityInst\":25.94,\"humidityMin\":24.73,\"humidityMax\":27.9,\"atmosphericPressureInst\":25.94,\"atmosphericPressureMin\":24.73,\"atmosphericPressureMax\":27.9,\"solarIrradianceInst\":25.94,\"solarIrradianceMin\":24.73,\"solarIrradianceMax\":27.9,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":2.52,\"windDirectionInst\":143.01,\"windSpeedGust\":5.43,\"windDirectionGust\":209.53}}"
        },
        {
            "key": "w/A327734/2023-10-21 01:00:00",
            "created": "2023-10-23T03:03:00.022Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":25.7,\"temperatureMin\":23.41,\"temperatureMax\":28.09,\"humidityInst\":81.09,\"humidityMin\":63.65,\"humidityMax\":90.11,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.21,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":1.44,\"windDirectionInst\":87.96,\"windSpeedGust\":7.24,\"windDirectionGust\":152.22}}"
        },
        {
            "key": "w/A327734/2023-10-21 01:10:53",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":81.09,\"temperatureMin\":63.65,\"temperatureMax\":90.11,\"humidityInst\":81.09,\"humidityMin\":63.65,\"humidityMax\":90.11,\"atmosphericPressureInst\":81.09,\"atmosphericPressureMin\":63.65,\"atmosphericPressureMax\":90.11,\"solarIrradianceInst\":81.09,\"solarIrradianceMin\":63.65,\"solarIrradianceMax\":90.11,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":1.44,\"windDirectionInst\":87.96,\"windSpeedGust\":7.24,\"windDirectionGust\":152.22}}"
        },
        {
            "key": "w/A327734/2023-10-21 02:00:00",
            "created": "2023-10-23T03:03:00.022Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":25.83,\"temperatureMin\":23.9,\"temperatureMax\":27.01,\"humidityInst\":74.41,\"humidityMin\":67.29,\"humidityMax\":90.55,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":554.3,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.19,\"solarIrradiation\":0.01,\"rain\":0,\"windSpeedInst\":4.68,\"windDirectionInst\":142.85,\"windSpeedGust\":12.67,\"windDirectionGust\":164.61}}"
        },
        {
            "key": "w/A327734/2023-10-21 02:10:46",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":583.41,\"temperatureMin\":554.3,\"temperatureMax\":1265.6,\"humidityInst\":583.41,\"humidityMin\":554.3,\"humidityMax\":1265.6,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":554.3,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":583.41,\"solarIrradianceMin\":554.3,\"solarIrradianceMax\":1265.6,\"solarIrradiation\":0.01,\"rain\":0.01,\"windSpeedInst\":4.68,\"windDirectionInst\":142.85,\"windSpeedGust\":12.67,\"windDirectionGust\":164.61}}"
        },
        {
            "key": "w/A327734/2023-10-21 03:00:00",
            "created": "2023-10-23T03:03:00.022Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":25.76,\"temperatureMin\":23.29,\"temperatureMax\":26.72,\"humidityInst\":74.28,\"humidityMin\":68.53,\"humidityMax\":85.62,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":513.3,\"atmosphericPressureMax\":1279.4,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.19,\"solarIrradiation\":0.01,\"rain\":0,\"windSpeedInst\":5.17,\"windDirectionInst\":204.6,\"windSpeedGust\":10.86,\"windDirectionGust\":160.75}}"
        },
        {
            "key": "w/A327734/2023-10-21 03:11:21",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":0.01,\"temperatureMin\":0,\"temperatureMax\":0,\"humidityInst\":0.01,\"humidityMin\":0,\"humidityMax\":0,\"atmosphericPressureInst\":0.01,\"atmosphericPressureMin\":0,\"atmosphericPressureMax\":0,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0.01,\"rain\":0.01,\"windSpeedInst\":5.17,\"windDirectionInst\":204.6,\"windSpeedGust\":10.86,\"windDirectionGust\":160.75}}"
        },
        {
            "key": "w/A327734/2023-10-21 04:00:00",
            "created": "2023-10-23T03:03:00.022Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.94,\"temperatureMin\":21.96,\"temperatureMax\":26.54,\"humidityInst\":76.74,\"humidityMin\":69.6,\"humidityMax\":88.49,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1299.8,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.19,\"solarIrradiation\":0.01,\"rain\":0,\"windSpeedInst\":8.66,\"windDirectionInst\":169.9,\"windSpeedGust\":16.29,\"windDirectionGust\":184.48}}"
        },
        {
            "key": "w/A327734/2023-10-21 04:10:21",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":0.01,\"temperatureMin\":0,\"temperatureMax\":0,\"humidityInst\":0.01,\"humidityMin\":0,\"humidityMax\":0,\"atmosphericPressureInst\":0.01,\"atmosphericPressureMin\":0,\"atmosphericPressureMax\":0,\"solarIrradianceInst\":0.01,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.66,\"windDirectionInst\":169.9,\"windSpeedGust\":16.29,\"windDirectionGust\":184.48}}"
        }
    ],
    "status": "success"
}
```

Example: Query three areas with a range of dates
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327732,A327734,A327735&date=2023-10-20 00:00:00,2023-10-20 02:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327732/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.848Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":21.94,\"temperatureMin\":20.07,\"temperatureMax\":23.05,\"humidityInst\":86.57,\"humidityMin\":74.42,\"humidityMax\":95.28,\"atmosphericPressureInst\":869.27,\"atmosphericPressureMin\":616.8,\"atmosphericPressureMax\":1259,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.37,\"windDirectionInst\":78.17,\"windSpeedGust\":14.56,\"windDirectionGust\":103.27}}"
        },
        {
            "key": "w/A327732/2023-10-20 00:10:46",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":88.02,\"temperatureMin\":81.64,\"temperatureMax\":95.28,\"humidityInst\":88.02,\"humidityMin\":81.64,\"humidityMax\":95.28,\"atmosphericPressureInst\":88.02,\"atmosphericPressureMin\":81.64,\"atmosphericPressureMax\":95.28,\"solarIrradianceInst\":88.02,\"solarIrradianceMin\":81.64,\"solarIrradianceMax\":95.28,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.37,\"windDirectionInst\":78.17,\"windSpeedGust\":14.56,\"windDirectionGust\":103.27}}"
        },
        {
            "key": "w/A327732/2023-10-20 00:17:04",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":22.81,\"temperatureMin\":22.78,\"temperatureMax\":22.92,\"humidityInst\":22.81,\"humidityMin\":22.78,\"humidityMax\":22.92,\"atmosphericPressureInst\":22.81,\"atmosphericPressureMin\":22.78,\"atmosphericPressureMax\":22.92,\"rain\":0,\"windSpeedInst\":2.4,\"windDirectionInst\":132.71,\"windSpeedGust\":6.02,\"windDirectionGust\":173.05}}"
        },
        {
            "key": "w/A327732/2023-10-20 00:32:42",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":630.6,\"temperatureMin\":630.4,\"temperatureMax\":630.7,\"humidityInst\":630.6,\"humidityMin\":630.4,\"humidityMax\":630.7,\"atmosphericPressureInst\":630.6,\"atmosphericPressureMin\":630.4,\"atmosphericPressureMax\":630.7,\"rain\":0,\"windSpeedInst\":0.26,\"windDirectionInst\":138.97,\"windSpeedGust\":6.76,\"windDirectionGust\":201.88}}"
        },
        {
            "key": "w/A327732/2023-10-20 00:47:52",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":630.5,\"temperatureMin\":630.5,\"temperatureMax\":630.6,\"humidityInst\":630.5,\"humidityMin\":630.5,\"humidityMax\":630.6,\"atmosphericPressureInst\":630.5,\"atmosphericPressureMin\":630.5,\"atmosphericPressureMax\":630.6,\"rain\":0,\"windSpeedInst\":2.91,\"windDirectionInst\":168.45,\"windSpeedGust\":9.05,\"windDirectionGust\":189.31}}"
        },
        {
            "key": "w/A327732/2023-10-20 01:00:00",
            "created": "2023-10-22T03:03:22.848Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":21.81,\"temperatureMin\":20.29,\"temperatureMax\":22.99,\"humidityInst\":86.68,\"humidityMin\":74.48,\"humidityMax\":95.73,\"atmosphericPressureInst\":869.61,\"atmosphericPressureMin\":616.8,\"atmosphericPressureMax\":1205.5,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":15.43,\"windDirectionInst\":72.49,\"windSpeedGust\":18.2,\"windDirectionGust\":80.94}}"
        },
        {
            "key": "w/A327732/2023-10-20 01:10:48",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":888.55,\"temperatureMin\":616.8,\"temperatureMax\":1180.2,\"humidityInst\":888.55,\"humidityMin\":616.8,\"humidityMax\":1180.2,\"atmosphericPressureInst\":888.55,\"atmosphericPressureMin\":616.8,\"atmosphericPressureMax\":1180.2,\"solarIrradianceInst\":888.55,\"solarIrradianceMin\":616.8,\"solarIrradianceMax\":1180.2,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":15.43,\"windDirectionInst\":72.49,\"windSpeedGust\":18.2,\"windDirectionGust\":80.94}}"
        },
        {
            "key": "w/A327732/2023-10-20 01:17:35",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":80.6,\"temperatureMin\":79.71,\"temperatureMax\":80.88,\"humidityInst\":80.6,\"humidityMin\":79.71,\"humidityMax\":80.88,\"atmosphericPressureInst\":80.6,\"atmosphericPressureMin\":79.71,\"atmosphericPressureMax\":80.88,\"rain\":0,\"windSpeedInst\":0.82,\"windDirectionInst\":139.56,\"windSpeedGust\":3.73,\"windDirectionGust\":122.16}}"
        },
        {
            "key": "w/A327732/2023-10-20 01:32:29",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":80.69,\"temperatureMin\":80.39,\"temperatureMax\":81.15,\"humidityInst\":80.69,\"humidityMin\":80.39,\"humidityMax\":81.15,\"atmosphericPressureInst\":80.69,\"atmosphericPressureMin\":80.39,\"atmosphericPressureMax\":81.15,\"rain\":0,\"windSpeedInst\":0.11,\"windDirectionInst\":105.6,\"windSpeedGust\":2.37,\"windDirectionGust\":105.64}}"
        },
        {
            "key": "w/A327732/2023-10-20 01:48:19",
            "created": "2023-10-20T03:02:48.846Z",
            "value": "{\"lonlat\":[-47.775714,-19.16201],\"weatherData\":{\"temperatureInst\":629.8,\"temperatureMin\":629.6,\"temperatureMax\":629.9,\"humidityInst\":629.8,\"humidityMin\":629.6,\"humidityMax\":629.9,\"atmosphericPressureInst\":629.8,\"atmosphericPressureMin\":629.6,\"atmosphericPressureMax\":629.9,\"rain\":0,\"windSpeedInst\":0.86,\"windDirectionInst\":163.85,\"windSpeedGust\":5.59,\"windDirectionGust\":57.21}}"
        },
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 00:10:38",
            "created": "2023-10-20T03:02:48.85Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":583.41,\"temperatureMin\":583.4,\"temperatureMax\":1265.6,\"humidityInst\":583.41,\"humidityMin\":583.4,\"humidityMax\":1265.6,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":583.41,\"solarIrradianceMin\":583.4,\"solarIrradianceMax\":1265.6,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 01:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.5,\"temperatureMin\":24.56,\"temperatureMax\":32.1,\"humidityInst\":56.62,\"humidityMin\":34.57,\"humidityMax\":78.31,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":567.8,\"atmosphericPressureMax\":1299.8,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.31,\"windDirectionInst\":163.18,\"windSpeedGust\":19.91,\"windDirectionGust\":179.47}}"
        },
        {
            "key": "w/A327734/2023-10-20 01:11:19",
            "created": "2023-10-20T03:02:48.85Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":0,\"temperatureMin\":0,\"temperatureMax\":0,\"humidityInst\":0,\"humidityMin\":0,\"humidityMax\":0,\"atmosphericPressureInst\":0,\"atmosphericPressureMin\":0,\"atmosphericPressureMax\":0,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.31,\"windDirectionInst\":163.18,\"windSpeedGust\":19.91,\"windDirectionGust\":179.47}}"
        },
        {
            "key": "w/A327735/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.855Z",
            "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":29.62,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":46.1,\"humidityMin\":29.57,\"humidityMax\":79.32,\"atmosphericPressureInst\":932.73,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327735/2023-10-20 00:10:38",
            "created": "2023-10-20T03:02:48.855Z",
            "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":932.67,\"temperatureMin\":583.4,\"temperatureMax\":1265.6,\"humidityInst\":932.67,\"humidityMin\":583.4,\"humidityMax\":1265.6,\"atmosphericPressureInst\":932.67,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":932.67,\"solarIrradianceMin\":583.4,\"solarIrradianceMax\":1265.6,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327735/2023-10-20 01:00:00",
            "created": "2023-10-22T03:03:22.855Z",
            "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":28.1,\"temperatureMin\":24.56,\"temperatureMax\":32.1,\"humidityInst\":57.25,\"humidityMin\":34.57,\"humidityMax\":78.31,\"atmosphericPressureInst\":952.24,\"atmosphericPressureMin\":567.8,\"atmosphericPressureMax\":1299.8,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.31,\"windDirectionInst\":163.18,\"windSpeedGust\":19.91,\"windDirectionGust\":179.47}}"
        },
        {
            "key": "w/A327735/2023-10-20 01:11:19",
            "created": "2023-10-20T03:02:48.855Z",
            "value": "{\"lonlat\":[-50.588414,-17.753447],\"weatherData\":{\"temperatureInst\":0,\"temperatureMin\":0,\"temperatureMax\":0,\"humidityInst\":0,\"humidityMin\":0,\"humidityMax\":0,\"atmosphericPressureInst\":0,\"atmosphericPressureMin\":0,\"atmosphericPressureMax\":0,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.31,\"windDirectionInst\":163.18,\"windSpeedGust\":19.91,\"windDirectionGust\":179.47}}"
        }
    ],
    "status": "success"
}
```

### Read with optional parameters

The optional parameters that the user can apply are:

1. version
2. regexp
3. count

#### version

By default, the API will return only the most updated version of a certain key. However, the user might want to retrieve older versions as well.

The result of a default request for the area `A327734` and date `2023-10-20 00:00:00` is:

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20 00:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        }
    ],
    "status": "success"
}
```

Notice that the "created" key is related to the time when the data was inserted in the database. It means that the most updated data was inserted at 2023-10-22T03:03:22.854Z.

After applying version=3:

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20 00:00:00&version=3'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-21T03:03:44.74Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-20T03:02:48.85Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        }
    ],
    "status": "success"
}
```

In this example, older versions of the same key are also retrieved.

#### Regexp

It is possible to apply regular expressions to the keys we are searching. For example, if one wants to query the area A327734 and only the dates with even numbers in the hour and zero for the minutes and seconds, the regexp would be `.*[02468]:00:00`.

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20 00:00:00,2023-10-21 00:00:00&regexp=.*[02468]:00:00'
```

Response:

Status code: 200 OK
```json
{
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 02:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.19,\"temperatureMin\":23.7,\"temperatureMax\":29.42,\"humidityInst\":56.47,\"humidityMin\":43.28,\"humidityMax\":79.69,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1262.8,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.01,\"windDirectionInst\":152.7,\"windSpeedGust\":18.1,\"windDirectionGust\":169.36}}"
        },
        {
            "key": "w/A327734/2023-10-20 04:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":26.69,\"temperatureMin\":25.07,\"temperatureMax\":28.78,\"humidityInst\":63.51,\"humidityMin\":51.26,\"humidityMax\":74.53,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":10.6,\"windDirectionInst\":81.96,\"windSpeedGust\":28.96,\"windDirectionGust\":75.93}}"
        },
        {
            "key": "w/A327734/2023-10-20 06:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.5,\"temperatureMin\":21.71,\"temperatureMax\":26.63,\"humidityInst\":73.27,\"humidityMin\":60.87,\"humidityMax\":85.16,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":9.55,\"windDirectionInst\":190.06,\"windSpeedGust\":18.1,\"windDirectionGust\":194.06}}"
        },
        {
            "key": "w/A327734/2023-10-20 08:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.37,\"temperatureMin\":21.54,\"temperatureMax\":25.36,\"humidityInst\":74.11,\"humidityMin\":66.31,\"humidityMax\":89.19,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1353,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":7.54,\"windDirectionInst\":206.05,\"windSpeedGust\":14.48,\"windDirectionGust\":170.41}}"
        },
        {
            "key": "w/A327734/2023-10-20 10:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.8,\"temperatureMin\":21.89,\"temperatureMax\":26.55,\"humidityInst\":74.4,\"humidityMin\":63.57,\"humidityMax\":88.8,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":368.7,\"atmosphericPressureMax\":1355.6,\"solarIrradianceInst\":77.64,\"solarIrradianceMin\":3.2,\"solarIrradianceMax\":113.25,\"solarIrradiation\":37.42,\"rain\":0,\"windSpeedInst\":8.43,\"windDirectionInst\":169.62,\"windSpeedGust\":23.53,\"windDirectionGust\":161.63}}"
        },
        {
            "key": "w/A327734/2023-10-20 12:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.23,\"temperatureMin\":25.47,\"temperatureMax\":30.91,\"humidityInst\":66.05,\"humidityMin\":51.65,\"humidityMax\":81.4,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":440.6,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":306.56,\"solarIrradianceMin\":120.32,\"solarIrradianceMax\":733.16,\"solarIrradiation\":237.41,\"rain\":0,\"windSpeedInst\":15.57,\"windDirectionInst\":131.04,\"windSpeedGust\":28.96,\"windDirectionGust\":139.3}}"
        },
        {
            "key": "w/A327734/2023-10-20 14:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.99,\"temperatureMin\":28.79,\"temperatureMax\":34.42,\"humidityInst\":58.32,\"humidityMin\":43.52,\"humidityMax\":75.15,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":505.8,\"atmosphericPressureMax\":1369.4,\"solarIrradianceInst\":894.13,\"solarIrradianceMin\":186.12,\"solarIrradianceMax\":1002.37,\"solarIrradiation\":605.79,\"rain\":0,\"windSpeedInst\":11.33,\"windDirectionInst\":89.51,\"windSpeedGust\":27.15,\"windDirectionGust\":125.15}}"
        },
        {
            "key": "w/A327734/2023-10-20 16:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":33.31,\"temperatureMin\":31.19,\"temperatureMax\":37.16,\"humidityInst\":47.62,\"humidityMin\":36.79,\"humidityMax\":70.91,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":828.71,\"solarIrradianceMin\":79.35,\"solarIrradianceMax\":1136.08,\"solarIrradiation\":888,\"rain\":0,\"windSpeedInst\":5.37,\"windDirectionInst\":25.2,\"windSpeedGust\":21.72,\"windDirectionGust\":65.03}}"
        },
        {
            "key": "w/A327734/2023-10-20 18:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":36.1,\"temperatureMin\":33.08,\"temperatureMax\":38.17,\"humidityInst\":40.72,\"humidityMin\":31.82,\"humidityMax\":65.71,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1299.8,\"solarIrradianceInst\":641.35,\"solarIrradianceMin\":108.78,\"solarIrradianceMax\":991.05,\"solarIrradiation\":763.15,\"rain\":0,\"windSpeedInst\":7.65,\"windDirectionInst\":180.4,\"windSpeedGust\":16.29,\"windDirectionGust\":359.91}}"
        },
        {
            "key": "w/A327734/2023-10-20 20:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.34,\"temperatureMin\":28.58,\"temperatureMax\":37.65,\"humidityInst\":61.44,\"humidityMin\":32.26,\"humidityMax\":72.34,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":123.22,\"solarIrradianceMin\":35.72,\"solarIrradianceMax\":590.85,\"solarIrradiation\":180.42,\"rain\":0,\"windSpeedInst\":21.99,\"windDirectionInst\":154.3,\"windSpeedGust\":45.26,\"windDirectionGust\":199.16}}"
        },
        {
            "key": "w/A327734/2023-10-20 22:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":27.97,\"temperatureMin\":26.07,\"temperatureMax\":30.02,\"humidityInst\":67.35,\"humidityMin\":57.95,\"humidityMax\":80.42,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":32.98,\"solarIrradiation\":10.3,\"rain\":0,\"windSpeedInst\":2.98,\"windDirectionInst\":15.67,\"windSpeedGust\":23.53,\"windDirectionGust\":35.24}}"
        }
    ],
    "status": "success"
}
```

#### count

The count parameter will provide the number of keys retrieved in the request for any case. It must be included as `count=true`

Example:
```shell
curl 'http://localhost:7000/read/climate-data?type=w&area_id=A327734&date=2023-10-20 00:00:00,2023-10-21 00:00:00&regexp=.*[02468]:00:00&count=true'
```

Response:

Status code: 200 OK
```json
{
    "count": 12,
    "result": [
        {
            "key": "w/A327734/2023-10-20 00:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.79,\"temperatureMin\":23.81,\"temperatureMax\":32.85,\"humidityInst\":55.59,\"humidityMin\":29.57,\"humidityMax\":76.83,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":6.6,\"windDirectionInst\":179.7,\"windSpeedGust\":12.67,\"windDirectionGust\":172.88}}"
        },
        {
            "key": "w/A327734/2023-10-20 02:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.19,\"temperatureMin\":23.7,\"temperatureMax\":29.42,\"humidityInst\":56.47,\"humidityMin\":43.28,\"humidityMax\":79.69,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1262.8,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":8.01,\"windDirectionInst\":152.7,\"windSpeedGust\":18.1,\"windDirectionGust\":169.36}}"
        },
        {
            "key": "w/A327734/2023-10-20 04:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":26.69,\"temperatureMin\":25.07,\"temperatureMax\":28.78,\"humidityInst\":63.51,\"humidityMin\":51.26,\"humidityMax\":74.53,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":10.6,\"windDirectionInst\":81.96,\"windSpeedGust\":28.96,\"windDirectionGust\":75.93}}"
        },
        {
            "key": "w/A327734/2023-10-20 06:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.5,\"temperatureMin\":21.71,\"temperatureMax\":26.63,\"humidityInst\":73.27,\"humidityMin\":60.87,\"humidityMax\":85.16,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":9.55,\"windDirectionInst\":190.06,\"windSpeedGust\":18.1,\"windDirectionGust\":194.06}}"
        },
        {
            "key": "w/A327734/2023-10-20 08:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.37,\"temperatureMin\":21.54,\"temperatureMax\":25.36,\"humidityInst\":74.11,\"humidityMin\":66.31,\"humidityMax\":89.19,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1353,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":0.02,\"solarIrradiation\":0,\"rain\":0,\"windSpeedInst\":7.54,\"windDirectionInst\":206.05,\"windSpeedGust\":14.48,\"windDirectionGust\":170.41}}"
        },
        {
            "key": "w/A327734/2023-10-20 10:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":24.8,\"temperatureMin\":21.89,\"temperatureMax\":26.55,\"humidityInst\":74.4,\"humidityMin\":63.57,\"humidityMax\":88.8,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":368.7,\"atmosphericPressureMax\":1355.6,\"solarIrradianceInst\":77.64,\"solarIrradianceMin\":3.2,\"solarIrradianceMax\":113.25,\"solarIrradiation\":37.42,\"rain\":0,\"windSpeedInst\":8.43,\"windDirectionInst\":169.62,\"windSpeedGust\":23.53,\"windDirectionGust\":161.63}}"
        },
        {
            "key": "w/A327734/2023-10-20 12:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":28.23,\"temperatureMin\":25.47,\"temperatureMax\":30.91,\"humidityInst\":66.05,\"humidityMin\":51.65,\"humidityMax\":81.4,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":440.6,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":306.56,\"solarIrradianceMin\":120.32,\"solarIrradianceMax\":733.16,\"solarIrradiation\":237.41,\"rain\":0,\"windSpeedInst\":15.57,\"windDirectionInst\":131.04,\"windSpeedGust\":28.96,\"windDirectionGust\":139.3}}"
        },
        {
            "key": "w/A327734/2023-10-20 14:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.99,\"temperatureMin\":28.79,\"temperatureMax\":34.42,\"humidityInst\":58.32,\"humidityMin\":43.52,\"humidityMax\":75.15,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":505.8,\"atmosphericPressureMax\":1369.4,\"solarIrradianceInst\":894.13,\"solarIrradianceMin\":186.12,\"solarIrradianceMax\":1002.37,\"solarIrradiation\":605.79,\"rain\":0,\"windSpeedInst\":11.33,\"windDirectionInst\":89.51,\"windSpeedGust\":27.15,\"windDirectionGust\":125.15}}"
        },
        {
            "key": "w/A327734/2023-10-20 16:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":33.31,\"temperatureMin\":31.19,\"temperatureMax\":37.16,\"humidityInst\":47.62,\"humidityMin\":36.79,\"humidityMax\":70.91,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":828.71,\"solarIrradianceMin\":79.35,\"solarIrradianceMax\":1136.08,\"solarIrradiation\":888,\"rain\":0,\"windSpeedInst\":5.37,\"windDirectionInst\":25.2,\"windSpeedGust\":21.72,\"windDirectionGust\":65.03}}"
        },
        {
            "key": "w/A327734/2023-10-20 18:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":36.1,\"temperatureMin\":33.08,\"temperatureMax\":38.17,\"humidityInst\":40.72,\"humidityMin\":31.82,\"humidityMax\":65.71,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1299.8,\"solarIrradianceInst\":641.35,\"solarIrradianceMin\":108.78,\"solarIrradianceMax\":991.05,\"solarIrradiation\":763.15,\"rain\":0,\"windSpeedInst\":7.65,\"windDirectionInst\":180.4,\"windSpeedGust\":16.29,\"windDirectionGust\":359.91}}"
        },
        {
            "key": "w/A327734/2023-10-20 20:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":30.34,\"temperatureMin\":28.58,\"temperatureMax\":37.65,\"humidityInst\":61.44,\"humidityMin\":32.26,\"humidityMax\":72.34,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1265.6,\"solarIrradianceInst\":123.22,\"solarIrradianceMin\":35.72,\"solarIrradianceMax\":590.85,\"solarIrradiation\":180.42,\"rain\":0,\"windSpeedInst\":21.99,\"windDirectionInst\":154.3,\"windSpeedGust\":45.26,\"windDirectionGust\":199.16}}"
        },
        {
            "key": "w/A327734/2023-10-20 22:00:00",
            "created": "2023-10-22T03:03:22.854Z",
            "value": "{\"lonlat\":[-50.59667,-17.749189],\"weatherData\":{\"temperatureInst\":27.97,\"temperatureMin\":26.07,\"temperatureMax\":30.02,\"humidityInst\":67.35,\"humidityMin\":57.95,\"humidityMax\":80.42,\"atmosphericPressureInst\":583.41,\"atmosphericPressureMin\":583.4,\"atmosphericPressureMax\":1271,\"solarIrradianceInst\":0,\"solarIrradianceMin\":0,\"solarIrradianceMax\":32.98,\"solarIrradiation\":10.3,\"rain\":0,\"windSpeedInst\":2.98,\"windDirectionInst\":15.67,\"windSpeedGust\":23.53,\"windDirectionGust\":35.24}}"
        }
    ],
    "status": "success"
}
```
