import React, { useState, useEffect, useRef}from 'react';
import { GoogleMap, useLoadScript, Marker, Circle } from '@react-google-maps/api';
//import { formatRelative } from 'date-fns';
import axios from 'axios';
import Table from '@mui/material/Table';
import TableHead from '@mui/material/TableHead';
import TableBody from '@mui/material/TableBody';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import styled from '@emotion/styled';
import {
  Paper,
} from "@mui/material";
import { Link } from 'react-router-dom';
import "./Map.css";


const TableContainer = styled.div`
  margin-top: 20px;
  margin-left: 0;
  width: 100%;
`;

const handleRestaurantClick = (id) => {
  sessionStorage.removeItem('selectedRestaurantId');
  sessionStorage.setItem('selectedRestaurantId', id);
};

const baseUrl = process.env.REACT_APP_BACKEND_API_URL;
const libraries = ['places'];
const mapContainerStyle = {
  width: '50vw',
  height: '100vh',
  position: 'absolute', // Position the map absolutely
  top: 0,
  right: 0, // Align the map to the right side
};

/*
const center = {
  lat: 7.2905715, // default latitude
  lng: 80.6337262, // default longitude
};
*/

const Map =  () => {

  
  useEffect(() => {
    console.log("useEffect");
    if (navigator.geolocation) {
      console.log("getting current position");
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setCurrentPosition({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
        },
        (error) => {
          console.error('Error getting the current position:', error);
        }
      );
      sessionStorage.setItem('currentLat', 40.0150);
      sessionStorage.setItem('currentLng', -105.2705);    
      getRestaurants();
    } else {
      console.error('Geolocation is not supported by this browser.');
    }

  }, []);
  

  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: 'AIzaSyAUTHkAmXLIwl3DM4VQLgLXXUzTbsH-H9Y',
    libraries,
  });

   const [currentPosition, setCurrentPosition] = useState(null);
  const [clickedPosition, setClickedPosition] = useState(null);
  const [radius, setRadius] = useState(0);

  
  
  
 //const [locations, setLocations] = useState([]);
//   // const [restarantData, setRestarantData] = useState([{}]);

   const [data, setData] = useState([]);

//   // const [data, setData] = useState(() => {
//   //   // Retrieve data from sessionStorage if available
//   //   const savedData = sessionStorage.getItem('savedData');
//   //   console.log(`saved data is ${savedData}`);
//   //   return savedData ? JSON.parse(savedData) : null;
//   // });

//   // console.log(` data is ${data}`);
var locations = [];
if (data) {
    locations = data.map(item => ({
      lng: item.geojson.coordinates[0],
      lat: item.geojson.coordinates[1],
      name: item.name,
    }));
  }

  const getRestaurants = async (lat, lng) => {
    console.log("getting restaurants");
    console.log("sessionStorage", sessionStorage)
    let accessToken = sessionStorage.getItem("access_token");

    if (!lat) {
      lat = sessionStorage.getItem("currentLat");
    }
    if (!lng) {
      lng = sessionStorage.getItem("currentLng");
    }
    console.log("Lat is", lat);
    console.log("Lng is", lng);
    try {
      const response = await axios.get( baseUrl+"/restaurants", {
      headers: {
        'Authorization': `Bearer ${accessToken}`
      },
      params: {
        lat: lat,
        long: lng,
        // lat: 40.0150,
        // long: -105.2705,
        distance: radius,
      },
    });
    console.log("success illa");
    console.log(response.data.data); 
    sessionStorage.removeItem('savedData');
    setData(response.data.data);
    sessionStorage.setItem('savedData', JSON.stringify(response.data.data));
    //setRestarantData(response.data.data);
    //console.log(restarantData);
    console.log("restarantData");
    console.log(data);
    //console.log("locations");
    //console.log(transformedData);
    
    
  
    } catch (error) {
      console.error('Error:', error);
    }
  }
 
  const handleMapClick = (event) => {
    console.log("handleMapClick");  
    setClickedPosition(null);
    console.log("Clicked position is", clickedPosition);
    sessionStorage.removeItem('clickedLat');
    sessionStorage.removeItem('clickedLng');
    sessionStorage.setItem('clickedLat', event.latLng.lat());
    sessionStorage.setItem('clickedLng', event.latLng.lng());
    setClickedPosition({
      lat: event.latLng.lat(),
      lng: event.latLng.lng(),
    });
    getRestaurants(event.latLng.lat(), event.latLng.lng());
  };

  const markerIcon = {
    url: 'https://cdn3.iconfinder.com/data/icons/map-14/144/Map-10-512.png', // URL of the custom icon image
  };

  const pinIcon = {
    url: 'https://www.iconpacks.net/icons/2/free-location-pin-icon-2965-thumb.png', // URL of the custom icon image
  };

  

  const radiusOptions = {
    strokeColor: '#0080ff', // Blue color for the stroke
    strokeOpacity: 0.8,
    strokeWeight: 2,
    fillColor: '#0080ff', // Blue color for the fill
    fillOpacity: 0.35,
  };
  const handleRadiusChange = (event) => {
    const newRadius = parseInt(event.target.value, 10);
    setRadius(newRadius);
    getRestaurants(sessionStorage.getItem('clickedLat'), sessionStorage.getItem('clickedLng'));
  };

  if (loadError) {
    return <div>Error loading maps</div>;
  }

  if (!isLoaded) {
    return <div>Loading maps</div>;
  }
  //console.log("the clicked position is");
  //console.log(clickedPosition.lat);
  //console.log("the current position is");
  //console.log(currentPosition.lat);

  return (
    <div style={{ display: 'flex' }}>
      <div style={{ flex: 1, position: 'relative' }}>
        <GoogleMap
          mapContainerStyle={mapContainerStyle}
          zoom={15} // You can adjust the zoom level as needed
          center={currentPosition} // Set the center to the user's current position
          onClick={handleMapClick}
        >
          {currentPosition && <Marker position={currentPosition} icon={{
            ...markerIcon,
            scaledSize: new window.google.maps.Size(50, 50),}}/>}
          {clickedPosition && (
          <>
            <Marker position={clickedPosition} icon={{
            ...pinIcon,
            scaledSize: new window.google.maps.Size(50, 50),}}/>
            <Circle
              center={clickedPosition}
              radius={radius} // Radius in meters
              options={radiusOptions}
            />
          </>
        )}

  {locations.map((location,index) => (
    console.log(location.name),
            <Marker
              key={index}
              position={{ lat: location.lat, lng: location.lng }}
              title={location.name}
              icon={{
                url: 'https://maps.google.com/mapfiles/ms/icons/blue-dot.png', // Default Google Maps marker icon
                labelOrigin: { x: 12, y: -15 }, // Position of the label relative to the icon
                scaledSize: new window.google.maps.Size(30, 30), // Size of the icon
                label: {
                  text: location.name,
                  color: 'black', // Label text color
                  fontSize: '14px', // Label text font size
                  fontWeight: 'bold' // Label text fo
                },
              }}
            />
          ))}

        
      </GoogleMap>


      </div>
      <div style={{ position: 'absolute', top: 10, left: 10 }}>
        {/* Input field to change the radius */}
        <input
          type="range"
          min="20"
          max="10000"
          value={radius}
          onChange={handleRadiusChange}
          step="10"
          style={{ width: '600px' }}
        />{' '}
        meters
      </div>
      <div style={{ position: 'absolute', top: 40, left: 10 }}>
      <TableContainer component={Paper} style={{ maxHeight: 1000, overflow: 'scroll' }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Restaurant Name</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
  {data ? (
    data.map((item) => (
      <TableRow key={item.id} className='restaurant-cell'>
        <TableCell >
          <Link to={`/review`} onClick={() => handleRestaurantClick(item.id)}>{item.name}</Link>
          <div>
            {/* Additional information like reviews, rating, etc. */}
            Category: {item.category}<br />
            Address: {item.full_address}<br />
            Score: {item.score}<br />
            Ratings: {item.ratings}<br />
            Price: {item.price_range}<br />
            {/* Add more information as needed */}
          </div>
        </TableCell>
      </TableRow>
    ))
  ) : (
    <TableRow>
      <TableCell colSpan={1}>No restaurants available</TableCell>
    </TableRow>
  )}
</TableBody>
        </Table>
      </TableContainer>
      </div>
    </div>
  );
};


export default Map;