import React from "react";
import API from "./api";
import moment from "moment";
import {CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis} from "recharts";

export class Graph extends React.Component {

    render() {
        return (
            <div className="graph">
                <div>
                    <LineChart width={500} height={300} data={this.props.weights}>
                        <Line type="monotone" dataKey="value" stroke="#8884d8"/>
                        <CartesianGrid strokeDasharray="3 3"/>
                        <XAxis dataKey="date" type="number"
                               tickFormatter={(unixTime) => moment.unix(unixTime).format('DD-MM-YY')}
                               domain={['auto', 'auto']}/>
                        <YAxis/>
                        <Tooltip labelFormatter={(unixTime) => moment.unix(unixTime).format('DD-MM-YY')}/>
                    </LineChart>
                </div>
            </div>
        );
    }
}