/* eslint react/prop-types: 0 */
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
        {trams.map((tram, i) => (
          <tr key={i}>
            <td className="py-1">
              <div className="tram-nb-bg bg-black text-white rounded flex items-center justify-center" style={{ width: "2rem", height: "2rem" }}>
                <span className="tram-nb" style={{ fontSize: "1rem" }}>{tram.line}</span>
              </div>
            </td>
            <td className="py-1">{tram.direction}</td>
            <td className="py-1">{tram.eta}</td>
          </tr>
        ))}

      </tbody>
    </table>
  )
}

export default TramTable;
