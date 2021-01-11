import React from 'react';
import ReactDOM from 'react-dom';
import {CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis} from 'recharts';
import './index.css';
import moment from 'moment'
import API from './api';
import DatePicker from "react-datepicker";
import {Button, Col, Form, FormControl, InputGroup} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "react-datepicker/dist/react-datepicker.css";
import 'react-datepicker/dist/react-datepicker-cssmodules.css';

class Graph extends React.Component {
    state = {data: []}

    componentDidMount() {
        API.get(`weights`)
            .then(res => {
                // weights are assumed to be fetched in order by date
                // convert ISO time from json to UnixTime
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
        this.state = {value: '', date: new Date()};
    }

    handleChange = (event) => {
        this.setState({value: Number(event.target.value)});
    }

    handleSubmit = (event) => {
        event.preventDefault();

        API.post(`weights`, {value: this.state.value, date: moment([this.state.date.getFullYear(), this.state.date.getMonth(), this.state.date.getDate()]).toISOString()})
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
                                <DatePicker selected={this.state.date} onChange={date => this.setState({date: date})} />
                            </Col>
                            <Col xs="auto">
                                <InputGroup>
                                    <Form.Control type="number" value={this.state.value} onChange={this.handleChange} placeholder="Enter weight value"/>

                                </InputGroup>
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
