import React from 'react';
import ReactDOM from 'react-dom';
import { LineChart, Line, CartesianGrid, XAxis, YAxis } from 'recharts';
import './index.css';
import axios from 'axios';
import moment from 'moment'
import API from './api';

//const data = [{weight: 56, timestamp:0}, {weight: 63, timestamp:1}, {weight: 58, timestamp: 3}];

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
                console.log(moment("0001-01-01T00:34:08+00:34").format('DD-MM-YY'))
            })
    }

    render() {
        return (
            <div className="graph">
                <div>
                    <LineChart width={400} height={400} data={this.state.data}>
                        <Line type="monotone" dataKey="value" stroke="#8884d8" />
                        <CartesianGrid stroke="#ccc" />
                        <XAxis dataKey="date" type="number" tickFormatter={(unixTime) => moment.unix(unixTime).format('DD-MM-YY')} domain = {['auto', 'auto']}/>
                        <YAxis />
                    </LineChart>
                </div>
            </div>
        );
    }
}

// ========================================

ReactDOM.render(
    <Graph />,
    document.getElementById('root')
);
