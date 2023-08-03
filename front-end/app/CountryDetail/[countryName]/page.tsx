import React from "react";

type Props = {
  countryName: string;
};

const CountryDetail = ({
  params: { countryName },
}: {
  params: { countryName: string };
}) => {
  return <div>{countryName}</div>;
};
export default CountryDetail;
