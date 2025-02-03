/* eslint react/prop-types: 0 */
import cn from 'classnames';

function TramTable({ trams }) {

  return (
    <table className="w-full text-left">
      <thead>
        <tr>
          <th className="pb-2">Linia</th>
          <th className="pb-2">Kierunek</th>
          <th className="pb-2">ETA</th>
        </tr>
      </thead>
      <tbody>
        {trams.map((tram) => (
          <tr key={tram.trip_id}>
            <td className="py-1">
              <div className="tram-nb-bg bg-white text-white rounded flex items-center justify-center" style={{ width: "2rem", height: "2rem" }}>
                <span className="tram-nb text-black" style={{ fontSize: "1.6rem" }}>{tram.line}</span>
              </div>
            </td>
            <td className="py-1">{tram.direction}
              {tram.isTripToCityCenter && (
                <i className="fa-solid fa-city ml-2"></i>
              )}
            </td>
            <td className={cn("py-1", tram.eta.split(" ")[0] <= 5 ? 'text-red-500' : '')}>{tram.eta}</td>
          </tr>
        ))}

      </tbody>
    </table>
  )
}

export default TramTable;
