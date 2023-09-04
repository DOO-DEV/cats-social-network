import {Meow} from "./meow.jsx";
import {api} from '../api'

export const Timeline = ({meows}) => {
    const onSubmit = (e) => {
        api.post("/meows", null, {
            params: {
                body: e.target.input
            }
        })
    }

    return (
        <div>
            <form onSubmit={onSubmit}>
                <div className="input-group">
                    <input  type="text" className="form-control" placeholder="What's happening?"/>
                        <div className="input-group-append">
                            <button className="btn btn-primary" type="submit">Meow</button>
                        </div>
                </div>
            </form>
            <div className="mt-4">
                {
                    meows.map(m => (
                        <Meow key={m.created_at} meow={m} />
                    ))
                }
            </div>
        </div>
    )
}