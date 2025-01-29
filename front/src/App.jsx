import { useEffect, useState } from 'react'

import Header from './components/Header'
import TramTable from './components/TramTable'

function App() {
  const [trams, setTrams] = useState([]);
  const [weather, setWeather] = useState({});

  useEffect(() => {
    fetch("http://localhost:8080/stop")
      .then(response => response.json())
      .then(data => {
        setTrams(data.trams);
        setWeather(data.weather);
      })
      .catch(error => console.error('Error fetching data:', error));

    const interval = setInterval(() => {
      fetch("http://localhost:8080/stop")
        .then(response => response.json())
        .then(data => {
          setTrams(data.trams);
        })
        .catch(error => console.error('Error fetching data:', error));
      }, 5000);

    return () => {
      clearInterval(interval);
    }
  }, [])

  return (
    <div className="bg-gray-800 shadow-lg rounded-lg p-4 main-view text-white">
      <Header weatherIcon={weather.icon} weatherVal={weather.temperature} />
      <TramTable trams={trams} />
    </div>
  )
}

export default App
