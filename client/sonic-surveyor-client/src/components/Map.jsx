import React, { useState, useEffect } from 'react';
import ReactMapGL, { Marker, Source, Layer } from 'react-map-gl';
import currentUserLocationIcon from "../assets/currentLocationIcon.png";
import RecordSoundButton from "./RecordSoundButton";
import GeolocationPermission from './GeolocationPermission';
import Loading from './Loading';
import proj4 from 'proj4';


const MAPBOX_TOKEN = process.env.REACT_APP_MAPBOX_TOKEN;
const BACKEND_API_IP = "http://172.20.10.4:26001";

const layerStyle = {
  id: 'polygon-layer',
  type: 'fill',
  paint: {
    'fill-color': [
      'match',
      ['get', 'ISOLVL'],
      0, 'transparent', // Set color to transparent when value is 0, 
      1, '#a0bbbf',
      2, '#b8d6d1',
      3, '#cfe4cc',
      4, '#e3f2bf',
      5, '#f4c683',
      6, '#e87d4d',
      7, '#cd463f',
      8, '#a11a4d',
      9, '#75095d',
      10, '#430a4a',
      '#ffffff' // Default color if no match found
    ],
    'fill-opacity': 0.95
  }
};
  
let geojsonData;


const Map = () => {
  const [viewport, setViewport] = useState({
    latitude: 0,
    longitude: 0,
    zoom: 1,
  });
  const [userLocation, setUserLocation] = useState(null);
  const [averageNoiseDecibel, setAverageNoiseDecibel] = useState(null);
  const [currentLocation, setCurrentLocation] = useState({
    latitude: null,
    longitude: null
  });
  const [loading, setLoading] = useState(true);
  const [loadingInfo, setLoadingInfo] = useState("");
  




  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;
          setViewport((prevState) => ({
            ...prevState,
            latitude,
            longitude,
          }));

          //Update current location value
          setCurrentLocation({
            latitude,
            longitude
          });

          // Fly to the user's current location
          flyToLocation(latitude, longitude);

          setUserLocation({ latitude, longitude });
        },
        (error) => {
          console.log(error);
          console.log(currentLocation);
        }
      );
    } else {
      console.log('Geolocation is not supported by this browser.');
    }



    setLoading(false);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);



  const handleUpdateLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;

          setViewport((prevViewport) => ({
            ...prevViewport,
            latitude,
            longitude
          }));

          setUserLocation({ latitude, longitude });
        },
        (error) => {
          console.log(error);
        }
      );
    } else {
      console.log('Geolocation is not supported by this browser.');
    }
  };

  const flyToLocation = (latitude, longitude) => {
    const startZoom = 1;
    const endZoom = 17;
    const duration = 2500; // milliseconds


    const zoomDiff = endZoom - startZoom;
    const latDiff = latitude - viewport.latitude;
    const lngDiff = longitude - viewport.longitude;

    let startTime = null;

    const updateViewport = (currentTime) => {
      if (!startTime) {
        startTime = currentTime;
      }

      const elapsedTime = currentTime - startTime;

      if (elapsedTime >= duration) {
        setViewport((prevState) => ({
          ...prevState,
          latitude,
          longitude,
          zoom: endZoom
        }));
        return;
      }

      const progress = elapsedTime / duration;
      const easedProgress = ease(progress);

      const newZoom = startZoom + zoomDiff * easedProgress;
      const newLatitude = viewport.latitude + latDiff * easedProgress;
      const newLongitude = viewport.longitude + lngDiff * easedProgress;

      setViewport((prevState) => ({
        ...prevState,
        zoom: newZoom,
        latitude: newLatitude,
        longitude: newLongitude
      }));

      requestAnimationFrame(updateViewport);
    };

    const ease = (t) => {
      return t < 0.5 ? 2 * t * t : -1 + (4 - 2 * t) * t;
    };

    requestAnimationFrame(updateViewport);
  };

  const handleAverageNoiseDecibel = (data) => {
    setAverageNoiseDecibel(data);
  };



  //Create the Noise Map
  const handleCreateNoiseMap = async () => {
    setLoading(true); // Set loading to true before making the API call

    try {
      const deleteResponse = await fetch(BACKEND_API_IP+'/cleanDatabase', {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (deleteResponse.ok) {
      setLoadingInfo ((await deleteResponse.text()).split('\n')[0]); 

      //Prepare and Format
      // const location = convertToEPSG3826(currentLocation.latitude, currentLocation.longitude);
      const location = convertToEPSG3826(25.043605, 121.533290);
      const latitudeValue = location.coordinates[0]; const longitudeValue = location.coordinates[1];
      const LWD500Value = averageNoiseDecibel;
      
      //Add to point-source.geojson
      const point_source = {
        "fileContents": "{\"type\":\"FeatureCollection\",\"name\":\"Point_Source\",\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"urn:ogc:def:crs:EPSG::3826\"}},\"features\":[{\"type\":\"Feature\",\"properties\":{\"PK\":1,\"LWD500\":"+LWD500Value+"},\"geometry\":{\"type\":\"Point\",\"coordinates\":["+latitudeValue+","+longitudeValue+",0.0]}}]}",
        "newFileName": "POINT_SOURCE.geojson",
        "generatedTableName": "POINT_SOURCE"
      }

      //Upload to Server
      const postResponse = await fetch(BACKEND_API_IP+'/importFileFromJSON', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(point_source),
        });

        if (postResponse.ok) {
          let info = (await postResponse.text()).split('\n')[0]
          setLoadingInfo (info); 

          //Final request
          const formData = new FormData();
          // ⭐️ TODO: Add Temp & Humidity from Weather APIs Later
          formData.append('temperature', '');
          formData.append('humidity', '');

          const mapCreationPostRequest = await fetch(BACKEND_API_IP+'/noiseLevelFromSourceOnlyNTUTArea', {
            method: 'POST',
            body: formData
          });

          if(mapCreationPostRequest.ok) {
            info = (await mapCreationPostRequest.text()).split('\n')[0];
            setLoadingInfo (info); 

            //Get the Contour Map Data from Backend
            const tableName = "CONTOURING_NOISE_MAP"


            const responseContourMap = await fetch(BACKEND_API_IP+`/exportTable?tableName=${tableName}`, {
              method: "GET",
            });
            if (responseContourMap.ok) {
            const response = await responseContourMap.json()
            geojsonData = convertFinalGeoJsonData(response);
            }
            
          }
        }
    }
    } catch (error) {
      // Handle network error or exception
      console.error('Error creating Sound Map:', error);
    } 
    finally {
      setLoading(false); // Set loading to false after the API call completes
    }
  };

  function convertToEPSG3826(latitude, longitude) {
    // Define the source and target coordinate reference systems
    const sourceCRS = 'EPSG:4326';  // WGS84 (latitude and longitude)
    const targetCRS = '+proj=tmerc +lat_0=0 +lon_0=121 +k=0.9999 +x_0=250000 +y_0=0 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs';  // Custom EPSG:3826

    // Import the proj4 library
    const proji4 = require('proj4').default;

  
    // Define the transformation function
    const transformFunction = proji4(sourceCRS, targetCRS).forward;
  
    // Convert the latitude and longitude to EPSG::3826
    const [x, y] = transformFunction([longitude, latitude]);
  
    // Create the GeoJSON point using the EPSG::3826 coordinates
    const point = {
      type: 'Point',
      coordinates: [x, y]
    };
  
    return point;
  }

  function convertFromEPSG3826ToWGS84(x, y) {
    const sourceProjection = 'EPSG:3826';
    const destinationProjection = 'EPSG:4326'; // WGS84 projection
  
    // Load the projection definitions
    proj4.defs(sourceProjection, '+proj=tmerc +lat_0=0 +lon_0=121 +k=0.9999 +x_0=250000 +y_0=0 +ellps=GRS80 +towgs84=0,0,0,0,0,0,0 +units=m +no_defs');
    proj4.defs(destinationProjection, '+proj=longlat +datum=WGS84 +no_defs');
  
    // Convert the point from source to destination projection
    const convertedPoint = proj4(sourceProjection, destinationProjection, [x, y]).map(coord => parseFloat(coord.toFixed(8)));
  
    // Return the converted point
    return convertedPoint;
  }
  
  function convertFinalGeoJsonData(geojsonData) {
    const modifiedGeoJsonData = { ...geojsonData };
  
    // Remove the 'crs' property
    delete modifiedGeoJsonData.crs;
  
    // Convert coordinates from EPSG:3826 to WGS84 (EPSG:4326)
    modifiedGeoJsonData.features.forEach(feature => {
      const coordinates = feature.geometry.coordinates;
      const convertedCoordinates = coordinates.map(polygon =>
        polygon.map(point => convertFromEPSG3826ToWGS84(point[0], point[1]))
      );
      feature.geometry.coordinates = convertedCoordinates;
    });
  
    return modifiedGeoJsonData;
  }


  return (
    <React.Fragment>
      {loading ?  
      <div>
        <Loading /> 
        {loadingInfo}
      </div> : <div>
      <div style={{ height: '100vh' }}>
        <ReactMapGL
          {...viewport}
          style={{width: '100vw', height: '100vh'}}
          mapStyle="mapbox://styles/mapbox/streets-v9"
          onMove={evt => setViewport(evt.viewport)}
          mapboxAccessToken={MAPBOX_TOKEN}
        >
          {userLocation && (
            <Marker longitude={userLocation.longitude} latitude={userLocation.latitude} color="red" />
          )}

       {
        geojsonData && (
        <Source type="geojson" data={geojsonData}>
          <Layer {...layerStyle} />
        </Source>
        )
       } 
        </ReactMapGL>

        <div className='informationArea'>
          <GeolocationPermission />
          {averageNoiseDecibel && <p>Average Noise Decibel: <b>{averageNoiseDecibel}</b></p>}
        </div>
        <div className='createNoiseMapArea'>
          {averageNoiseDecibel && <button className='createNoiseMapButton' onClick={handleCreateNoiseMap}><b>Create Noise Map</b></button>}
        </div>
        <div className = "additionalButtons">
            <button onClick={handleUpdateLocation} className='currentLocationButton'>
              <img src={currentUserLocationIcon} alt="Current User Location Icon" width="48" height="48"/>
            </button>
            <RecordSoundButton onSendAverageNoiseDecibel={handleAverageNoiseDecibel} />
        </div>
      </div>
      </div>}
    </React.Fragment>
  );
};

export default Map;
