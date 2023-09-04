import {Search} from "./components/search.jsx";
import {Timeline} from "./components/timeline.jsx";
import {api} from "./api/index.js";
import {useEffect} from "react";

function App() {
  return (
      <div className="container py-5">
        <div className="row mb-4">
          <h1 className="col-12">Meower</h1>
        </div>
        <div className="row">
            <div className="col">
                <Timeline  meows={[]}/>
            </div>
            <div className="col">
                <Search />
            </div>
        </div>
      </div>
  )
}

export default App
