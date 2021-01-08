import React from 'react';
import ReactDOM from 'react-dom';
import {CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis} from 'recharts';
import './index.css';
import moment from 'moment'
import API from './api';
import {Button, Col, Form} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';

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
                        <Legend/>
                    </LineChart>
                </div>
            </div>
        );
    }
}

class WeightForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: ''};
    }

    handleChange = (event) => {
        this.setState({value: Number(event.target.value)});
    }

    handleSubmit = (event) => {
        alert('A weight was submitted: ' + this.state.value);
        event.preventDefault();
        API.post(`weights`, {value: this.state.value})
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
    }

    render() {
        return (
            <>
                <form onSubmit={this.handleSubmit}>
                    <Form.Group>
                        <Form.Row className="align-items-center">
                            <Col xs="auto">
                                <Form.Label column>Enter Weight Value</Form.Label>
                            </Col>
                            <Col xs="auto">
                                <Form.Control type="number" value={this.state.value} onChange={this.handleChange} placeholder="Enter weight value"/>
                            </Col>
                            <Button variant="primary" type="submit">
                                Submit
                            </Button>
                        </Form.Row>
                    </Form.Group>
                </form>
            </>
        );
    }
}

class Page extends React.Component {
    render() {
        return (
            <>
                <Graph/>
                <WeightForm/>
            </>
        )
    }
}

// ========================================

ReactDOM.render(
    <Page/>,
    document.getElementById('root')
);
