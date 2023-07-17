# url-shortener

A very simple url shortener written in golang.

## Install
### Docker compose
Rename `example.env` to `.env`
```
mv example.env .env
```
Set variables in `.env`
Run containers with `docker-compose`
```
docker-compose up -d --build
```

## Usage
### Shorten URL
Shorten a URL

**URL** : `/shorten`

**Method** : `POST`

**Data constraints**

Provide original URL

```json
{
    "url": "example.com"
}
```
#### Success Response

**Condition** : If everything is OK sending back shortened url.

**Code** : `200 OK`

**Content example**

```json
{
    "url": "http://testserver/3"
}
```
### Redirect short URL
Redirect short URL to original URL

**URL** : `/{shortcode}`

**Method** : `GET`
#### Success Response

**Condition** : If everything is OK redirecting.

**Code** : `303 See Other`
### Info about a shortened URL
Getting information about a shortened URL

**URL** : `/info/{shortcode}`

**Method** : `GET`
#### Success Response

**Condition** : If everything is OK sending back information about short URL.

**Code** : `200 OK`

**Content example**

```json
{
    "shortcode": 2,
    "url": "https://example.com",
    "active": true,
    "created": "2023-07-17 12:48:28"
}
```
### Turning short URL on and off
Change the active state of the short URL.

**URL** : `/{shortcode}`

**Method** : `PUT`
#### Success Response

**Condition** : If everything is OK toggling short URL active state.

**Code** : `200 OK`

### Permanently delete short URL
Delete short URL from database.

**URL** : `/{shortcode}`

**Method** : `DELETE`
#### Success Response

**Condition** : If everything is OK deleting URL from database.

**Code** : `200 OK`


# Licence 
This module is open-sourced software licensed under the [GPL-3.0 license](https://opensource.org/license/gpl-3-0/)