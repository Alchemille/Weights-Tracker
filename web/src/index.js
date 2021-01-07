import React from 'react';
import ReactDOM from 'react-dom';
import {CartesianGrid, Line, LineChart, XAxis, YAxis, Tooltip, Legend} from 'recharts';
import './index.css';
import moment from 'moment'
import API from './api';

class Graph extends React.Component {
    state = {data: []}

    componentDidMount() {
        API.get(`weights`)
            .then(res => {

                const weights = res.data.map((w) => {
                    return {...w, date: moment(w.date).unix()};
                });
                this.setState({data: weights});
                console.log(weights);
            })
    }

    render() {
        return (
            <div className="graph">
                <div>
                    <LineChart width={500} height={300} data={this.state.data}>
                        <Line type="monotone" dataKey="value" stroke="#8884d8"/>
                        <CartesianGrid strokeDasharray="3 3"/>
                        <XAxis dataKey="date" type="number"
                               tickFormatter={(unixTime) => moment.unix(unixTime).format('DD-MM-YY')}
                               domain={['auto', 'auto']}/>
                        <YAxis/>
                        <Tooltip labelFormatter={(unixTime) => moment.unix(unixTime).format('DD-MM-YY')}/>
                        <Legend />
                    </LineChart>
                </div>
            </div>
        );
    }
}

// ========================================

ReactDOM.render(
    <Graph/>,
    document.getElementById('root')
);
