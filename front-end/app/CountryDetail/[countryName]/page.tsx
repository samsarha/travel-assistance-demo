import React from "react";
import Image from "next/image"

const CountryDetail = ({params: { countryName },}: {params: { countryName: string };}) => {
  return (
  <div className="flex flex-col bg-[#f1f5f9] min-h-screen">
    
    <h1 className="text-center font-extrabold text-[#222328 text-[32px]">{countryName}</h1>

    <div className="flex flex-row">

    <div className="pl-4 ml-4 block mb-12 w-full rounded-lg bg-white shadow-[0_2px_15px_-3px_rgba(0,0,0,0.07),0_10px_20px_-2px_rgba(0,0,0,0.04)] dark:bg-neutral-700">
      <h2 className="font-bold mb-4 mt-4">Country Info</h2>

      <div className="flex flex-row">
        <div className="w-1/2 pl-12">
        <h3 className="text-[32px]">GDP</h3>
        <p>12 million</p>
        </div>

        <div className="pl-12">
        <h3 className="text-[32px]">Exchange rate</h3>
        <p>1 NAR / 120 MZN</p>
        </div>

      </div>


    </div> 
    
    <div className="pl-4 ml-4 block mr-8 mb-12 min-w-[25rem] rounded-lg bg-white shadow-[0_2px_15px_-3px_rgba(0,0,0,0.07),0_10px_20px_-2px_rgba(0,0,0,0.04)] dark:bg-neutral-700">
      <h2 className="font-bold mb-4 mt-4">Weather forecast</h2>
        
        <div className="flex mb-2 items-center">
          <p className="mr-2">Thursday, August 03</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> 12 / 32 °C (Rainy)</p>
          </div>

          <div className="flex mb-2 items-center">
          <p className="mr-2">Friday, August 04</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> 12 / 32 °C (Rainy)</p>
          </div>

          <div className="flex mb-2 items-center">
          <p className="mr-2">Saturday, August 05</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> 12 / 32 °C (Rainy)</p>
          </div>

          <div className="flex mb-2 items-center">
          <p className="mr-2">Sunday, August 06</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> 12 / 32 °C (Rainy)</p>
          </div>

          <div className="flex mb-2 items-center">
          <p className="mr-2">Monday, August 07</p>
          <Image src="https://openweathermap.org/img/wn/10d@2x.png" width={42} height={42} alt="weather"/>
          <p className="text-[14px]"> 12 / 32 °C (Rainy)</p>
          </div>
    </div>
    </div>

     

  </div>);
};
export default CountryDetail;
