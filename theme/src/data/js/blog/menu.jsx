import React from 'react'


export default class MenuButton extends React.Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <div className="menu-button" onClick={this.props.toggleSearchContainer}>
        <span></span>
        <span></span>
        <span></span>
      </div>
    )
  }
}
