// import React, { useEffect, useState } from 'react';

// const GeolocationPermission = () => {
//   const [permissionGranted, setPermissionGranted] = useState(false);

//   useEffect(() => {
//     const requestPermission = async () => {
//       if (navigator.geolocation) {
//         const permissionStatus = await navigator.permissions.query({ name: 'geolocation' });
//         setPermissionGranted(permissionStatus.state === 'granted');
//       }
//     };

//     requestPermission();
//   }, []);

//   return (
//     <div>
//       {permissionGranted ? (
//         "" //Permission granted.
//       ) : (
//         <div> 
//             <p>Requesting geolocation permission...</p>
//             <p>Geolocation permission is required for this app.</p>
//             <p>Please enable geolocation in your browser settings.</p>
//         </div>
//       )}
//     </div>
//   );
// };

// export default GeolocationPermission;


import React, { useEffect, useState } from 'react';

const GeolocationPermission = () => {
  const [permissionGranted, setPermissionGranted] = useState(false);

  useEffect(() => {
    const requestPermission = () => {
      if ('geolocation' in navigator) {
        navigator.permissions.query({ name: 'geolocation' }).then(permissionStatus => {
          if (permissionStatus.state === 'granted') {
            setPermissionGranted(true);
          } else if (permissionStatus.state === 'prompt') {
            navigator.geolocation.getCurrentPosition(() => {
              setPermissionGranted(true);
            }, () => {
              setPermissionGranted(false);
            });
          } else {
            setPermissionGranted(false);
          }
        });
      } else {
        setPermissionGranted(false);
      }
    };

    requestPermission();
  }, []);

  return (
    <div>
      {permissionGranted ? (
        ""//Geolocation permission granted.
      ) : (
        <div> 
             <p>Requesting geolocation permission...</p>
             <p>Geolocation permission is required for this app.</p>
             <p>Please enable geolocation in your browser settings.</p>
         </div>
      )}
    </div>
  );
};

export default GeolocationPermission;
