/* eslint react/prop-types: 0 */
import { useState, useEffect } from 'react';

function Header({weatherVal, weatherIcon}) {
    const [currentTime, setCurrentTime] = useState(new Date());

    useEffect(() => {
        const intervalId = setInterval(() => {
            setCurrentTime(new Date());
        }, 1000);

        return () => clearInterval(intervalId);
    }, []);

    const formattedTime = currentTime.toLocaleTimeString();
    const formattedDate = currentTime.toLocaleDateString();

    return (
        <header className="flex justify-between items-center border-b pb-2 mb-4">
            <div className="flex justify-between items-center">
                <h1 className="text-2xl font-bold">CZYŻYNY</h1>
            </div>
            <div>
                <p className="text-2xl"><span>{formattedTime}</span></p>
            </div>
            <div className="flex items-center space-x-4">
                <p className="text-lg text-white-600"><span>{formattedDate}</span></p>
                <div className="flex items-center">
                    <img style={{ height: "2rem" }} src={`https://openweathermap.org/img/wn/${weatherIcon}@2x.png`} />
                    <span className="text-xl font-semibold">{weatherVal}°C</span>
                </div>
            </div>
        </header>
    );
}

export default Header;
