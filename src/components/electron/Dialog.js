import React, { Component } from 'react'

const dialogOptions = {
  type: 'info',
  title: 'Information',
  message: "This is an information dialog. Isn't it nice?",
  buttons: ['Yes', 'No']
}

export default class App extends Component {
  handleClick = () => {
    window.remote.dialog.showMessageBox(dialogOptions, (index) => {
      console.log('information-dialog-selection', index)
    })
  }

  render() {
    return (
      <div>
        <button onClick={this.handleClick}>test</button>
        {this.props.children}
      </div>
    )
  }
}