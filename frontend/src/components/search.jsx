import {useState} from "react";
import {Meow} from "./meow.jsx";
import {api} from '../api'

export const Search = () => {
    const [meows, setMeows] = useState([])
    const [lastQuery, setLastQuery] = useState()

    const searchMeows = (query) => {
        if(query == lastQuery) {
            return
        }
        api.get(`/search`, {params: { query },}).then(({data}) => {
            setMeows(data)
            setLastQuery(query)
        }).catch(e => console.log(e))
    }

    const onKeyUp = ({target}) => {
        searchMeows(target.value)
    }
    return (
        <div>
            <input onKeyUp={onKeyUp}  type="text" className="form-control" placeholder="Search..."/>
            <div className="mt-4">
                {
                    meows.map(m =>
                        <Meow key={m.id} meow={m} />
                    )
                }
            </div>
        </div>
    )
}