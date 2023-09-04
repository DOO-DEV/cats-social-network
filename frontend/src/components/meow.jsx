import {format} from 'timeago.js'

export const Meow = ({meow}) => {
    return (
        <div className="card">
            <div className="card-body">
                <p className="card-text" dangerouslySetInnerHTML={{__html: meow.body}}></p>
                <p className="card-text">
                    <small className="text-muted">
                        {format(Date.parse(meow.created_at))}
                    </small>
                </p>
            </div>
        </div>

    )
}