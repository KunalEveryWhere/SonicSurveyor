import './App.css';
import Map from "./components/Map"
import 'mapbox-gl/dist/mapbox-gl.css';

function App() {
  return (
    <div className="App">
      <h1 className='header'><span className='spanColor'>Sonic</span>Surveyor<br/>
      <span className='spanSubtitle'>SIxSD, NTUT (2023)</span></h1>
      <Map />
    </div>
  );
}

export default App;
