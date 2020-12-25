import React from 'react';
import ReactDOM from 'react-dom';
import { LineChart, Line, CartesianGrid, XAxis, YAxis } from 'recharts';
import './index.css';
import axios from 'axios';
import API from './api';

//const data = [{weight: 56, timestamp:0}, {weight: 63, timestamp:1}, {weight: 58, timestamp: 3}];

class Graph extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data: [{weight: 56, timestamp:0}, {weight: 63, timestamp:1}, {weight: 58, timestamp: 3}],
        };
    }

    componentDidMount() {
        API.get(`weights`)
            .then(res => {
                const weights = res.data;
                //this.setState({ persons });
                console.log(weights)
            })
    }

    render() {
        return (
            <div className="graph">
                <div>
                    <LineChart width={400} height={400} data={this.state.data}>
                        <Line type="monotone" dataKey="weight" stroke="#8884d8" />
                        <CartesianGrid stroke="#ccc" />
                        <XAxis dataKey="timestamp" />
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
