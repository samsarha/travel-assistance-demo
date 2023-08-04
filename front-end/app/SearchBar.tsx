import React, { ChangeEvent, useState } from "react";
import { countryNames } from "./constants";
import Link from "next/link";

const SearchBar = () => {
  const [searchInput, setSearchInput] = useState("");

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setSearchInput(e.target.value);
  };

  const filteredCountries: {name: string, code:string}[] = countryNames.filter((obj) => {
    if (searchInput.length > 0) {
      return obj.name.toLowerCase().includes(searchInput.toLowerCase());
    }
  });

  return (
    <div className="flex flex-col items-center">
      <form className="max-w-sm px-4 flex">
        <div className="relative flex-grow">
          <input
            type="text"
            placeholder="Search country"
            onChange={handleChange}
            value={searchInput}
            className="w-96 py-3 pl-12 pr-4 border rounded-md outline-none bg-gray-50 focus:bg-white focus:border-indigo-600"
          />
        </div>

        <button className="px-8 text-white bg-indigo-600 border-l rounded ml-2">
          Search
        </button>
      </form>

      <div className="block w-full max-w-[18rem] rounded-lg bg-white shadow-[0_2px_15px_-3px_rgba(0,0,0,0.07),0_10px_20px_-2px_rgba(0,0,0,0.04)] dark:bg-neutral-700">
        <ul className="w-full">
          {filteredCountries.map((countryObj, index) => (
            <li
              className="w-full border-b-2 border-neutral-100 border-opacity-100 px-4 py-3 dark:border-opacity-50 hover:bg-indigo-100"
              key={index}
            >
              <Link href={`/CountryDetail/${countryObj.code}`}>{countryObj.name}</Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default SearchBar;
