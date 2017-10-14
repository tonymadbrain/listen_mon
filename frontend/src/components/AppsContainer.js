import React, { Component } from 'react'
import axios from 'axios'

class AppsContainer extends Component {
  componentDidMount() {
    axios.get('http://127.0.0.1:1234/apps.json')
         .then(response => {
            console.log(response)
            this.setState({apps: response.data})
          })
         .catch(error => console.log(error))
  }
  constructor(props) {
    super(props)
    this.state = {
      apps: []
    }
  }
  render() {
    return (
      <div>
        <ol>
        {this.state.apps.map((app) => {
          return(
            <li className="tile" key={app.port}>{app.app} - <a target="_blank" href={`http://127.0.0.1:${app.port}`}>{app.port}</a></li>
          )
        })}
        </ol>
      </div>
    );
  }
}

export default AppsContainer;
