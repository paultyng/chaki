import React, { Component } from 'react';

class Navbar extends Component {
  render() {
    return (
      <div className="navbar navbar-default">
        <div className="container">
          <div className="navbar-header">
            <a className="navbar-brand" href="/">Chaki</a>
          </div>

          {this.props.children}
        </div>
      </div>
    )
  }
}

export default Navbar;
