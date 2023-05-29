import React, { useState, useEffect } from 'react';
import ReactMapGL, { Marker } from 'react-map-gl';
import currentUserLocationIcon from "../assets/currentLocationIcon.png";
import RecordSoundButton from "./RecordSoundButton";

const MAPBOX_TOKEN = process.env.REACT_APP_MAPBOX_TOKEN;

const Map = () => {
  const [viewport, setViewport] = useState({
    latitude: 0,
    longitude: 0,
    zoom: 1,
  });
  const [userLocation, setUserLocation] = useState(null);
  const [averageNoiseDecibel, setAverageNoiseDecibel] = useState(null);

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
          // Fly to the user's current location
          flyToLocation(latitude, longitude);

          setUserLocation({ latitude, longitude });
        },
        (error) => {
          console.log(error);
        }
      );
    } else {
      console.log('Geolocation is not supported by this browser.');
    }
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
    const frameRate = 120; // frames per second
    const frameDuration = 1000 / frameRate; // milliseconds

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


  return (
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
      </ReactMapGL>
      <div className='informationArea'>

        {averageNoiseDecibel && <p>Average Noise Decibel: <b>{averageNoiseDecibel}</b></p>}
      </div>
      <div className = "additionalButtons">
          <button onClick={handleUpdateLocation} className='currentLocationButton'>
            <img src={currentUserLocationIcon} alt="Current User Location Icon" width="48" height="48"/>
          </button>
          <RecordSoundButton onSendAverageNoiseDecibel={handleAverageNoiseDecibel} />
      </div>
    </div>
  );
};

export default Map;
