"use client";

import React, { useState, useEffect } from "react";
import Image from "next/image"


interface Country {
  countryname: string
  gdp: number
  population: number
  exchangerate: number
  weather: Weather[]
}

interface Weather {
  countryname: string
  iconUri: string
  min: number
  max: number
  description: string
}

const CountryDetail = ({params: { countryCode },}: {params: { countryCode: string };}) => {


  const [country, setCountry] = useState<Country>();

  useEffect(()  => {
    const fetchCountry = async() => {
  
      try {
        const response = await fetch("http://localhost:8080/country/mz",{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      });

      if (response.ok) {
        const result = await response.json();
        setCountry(result)
      }

    }
      catch(error) {
        alert(error)
      }
    };

    fetchCountry();

  }, []);


  
  return (
  <div className="flex flex-col bg-[#f1f5f9] min-h-screen">
    
    <h1 className="text-center font-extrabold text-[#222328 text-[32px]">{country?.countryname}</h1>

    <div className="flex flex-row">

    <div className="pl-4 ml-4 block mb-12 w-full rounded-lg bg-white shadow-[0_2px_15px_-3px_rgba(0,0,0,0.07),0_10px_20px_-2px_rgba(0,0,0,0.04)] dark:bg-neutral-700">
      <h2 className="font-bold mb-4 mt-4">Country Info</h2>
      

      <div className="flex flex-row">

      <div className="w-1/3 pl-12">
        <h3 className="text-[32px]">Population</h3>
        <p>{country?.population}</p>
        </div>

        <div className="w-1/3 pl-12">
        <h3 className="text-[32px]">GDP</h3>
        <p>{country?.gdp} USD</p>
        </div>

        <div className="w-1/3 pl-12">
        <h3 className="text-[32px]">Exchange rate</h3>
        <p>1 Unit / {country?.exchangerate} USD</p>
        </div>

      </div>


    </div> 
    
    <div className="pl-4 ml-4 block mr-8 mb-12 min-w-[30rem] rounded-lg bg-white shadow-[0_2px_15px_-3px_rgba(0,0,0,0.07),0_10px_20px_-2px_rgba(0,0,0,0.04)] dark:bg-neutral-700">
      <h2 className="font-bold mb-4 mt-4">Weather forecast</h2>

      {country?.weather.map((weatherObj, index) => (

        <div className="flex mb-2 items-center">
          <p className="mr-2">{weatherObj.countryname}</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> {weatherObj.min} / {weatherObj.max} °C ({weatherObj.description})</p>
          </div>
          ))}

    </div>
    </div>     

  </div>);
};
export default CountryDetail;
