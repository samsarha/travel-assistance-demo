# travel-assistance-demo

A demo travel assistance web-app

The frontend for this project was built with React/NextJS + Tailwind and for the backend Go was used with the Gin Framework.

The external apis for this projects are:

 openweathermap.org -> To get the weather forecast
 
 exchangeratesapi.io -> To get the exchange rate
 
 worldbank.org -> To get the population and GDP

Aditionally Auth0 was used to handle user authentication (https://auth0.com/).

The project still runs in dev at http://localhost:3000/
The underlying Rest api endpoint is http://localhost:8080/country/{countrycode} in which {countryCode} represents the
ISO 3166-1 alpha-2 code of the queried country. Ex: http://localhost:8080/country/mz will fetch data for Mozambique.
