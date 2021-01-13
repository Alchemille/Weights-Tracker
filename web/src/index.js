import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import "react-datepicker/dist/react-datepicker.css";
import 'react-datepicker/dist/react-datepicker-cssmodules.css';
import {Graph} from "./Graph";
import {WeightForm} from "./WeightForm";
import API from "./api";
import moment from "moment";

class Page extends React.Component {
    state = {weights:[]}

    componentDidMount() {
        this.loadWeights();
    }

    render() {
        return (
            <>
                <Graph weights={this.state.weights}/>
                <WeightForm onSubmit={this.handleSubmit}/>
            </>
        )
    }

    handleSubmit = () => {
        this.loadWeights();
    }


    loadWeights() {
        API.get(`weights`)
            .then(res => {
                // weights are assumed to be fetched in order by date
                // convert ISO time from json to UnixTime
                const weights = res.data.map((w) => {
                    return {...w, date: moment(w.date).unix()};
                });
                this.setState({ weights: weights });
                console.log(weights);
            })
    }
}

// ========================================

ReactDOM.render(
    <Page/>,
    document.getElementById('root')
);
