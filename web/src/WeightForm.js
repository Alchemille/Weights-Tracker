import React from "react";
import API from "./api";
import moment from "moment";
import {Button, Col, Form, InputGroup} from "react-bootstrap";
import DatePicker from "react-datepicker";

export class WeightForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: '', date: new Date(), errors:{}};
    }

    validateForm = () => {
        let valid = true;
        let errors = {};

        if (!this.state.value || isNaN(this.state.value)) {
            errors["value"] = "Please enter a number"
            valid = false
        }

        if (this.state.date > new Date()) {
            errors["date"] = "Future dates are not valid"
            valid = false
        }
        this.setState({errors: errors});
        return valid
    }

    handleChange = (event) => {
        this.setState({value: Number(event.target.value)});
    }

    handleSubmit = (event) => {
        event.preventDefault();

        if (!this.validateForm()) {
            return
        }

        API.post(`weights`, {
            value: this.state.value,
            date: moment([this.state.date.getFullYear(), this.state.date.getMonth(), this.state.date.getDate()]).toISOString()
        })
            .then((response) => {
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
                                <span style={{color: "red"}}>{this.state.errors["date"]}</span>
                            </Col>
                            <Col xs="auto">
                                <InputGroup>
                                    <Form.Control type="number" value={this.state.value} onChange={this.handleChange}
                                                  placeholder="Enter weight value"/>


                                </InputGroup>
                            </Col>
                            <br/>
                            <span className='mr-2' style={{color: "red"}}>{this.state.errors["value"]}</span>
                            <br/>
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