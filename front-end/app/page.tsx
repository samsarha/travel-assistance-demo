"use client";

import SearchBar from "./SearchBar";

const Home = () => {
  return (
    <div className="flex flex-col items-center my-8">
      <h1 className="font-extrabold text-[#222328 text-[32px]">Travel Assistant</h1>
      <p className="mb-4">Find countries for your next trip</p>
      <SearchBar/>
    </div>
  )
}

export default Home;