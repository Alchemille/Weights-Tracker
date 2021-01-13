import React from "react";
import API from "./api";
import moment from "moment";
import {Button, Col, Form, InputGroup} from "react-bootstrap";
import DatePicker from "react-datepicker";

export class WeightForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: '', date: new Date()};
    }

    handleChange = (event) => {
        this.setState({value: Number(event.target.value)});
    }

    handleSubmit = (event) => {
        event.preventDefault();

        API.post(`weights`, {
            value: this.state.value,
            date: moment([this.state.date.getFullYear(), this.state.date.getMonth(), this.state.date.getDate()]).toISOString()
        })
            .then((response) => {
                console.log(response);
                this.props.onSubmit();
            })
            .catch(error => {
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
                                <DatePicker selected={this.state.date} onChange={date => this.setState({date: date})}/>
                            </Col>
                            <Col xs="auto">
                                <InputGroup>
                                    <Form.Control type="number" value={this.state.value} onChange={this.handleChange}
                                                  placeholder="Enter weight value"/>

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