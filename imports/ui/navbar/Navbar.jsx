import React, { Component } from 'react';

export default class Navbar extends Component {
  getButtons() {
    return [
      { _id: 'dashboard-nav-button', url: 'dashboard', target: '', name: 'Dashboard' },
      { _id: 'github-nav-button', url: 'https://github.com/brianshef', target: '_blank', name: 'GitHub' }
    ];
  }

  renderButtons() {
    return this.getButtons().map((button) => (
      <li key={button._id} id={button._id}><a href={button.url} target={button.target}>{button.name}</a></li>
    ));
  }

  render() {
    return (
      <div class="container-fluid" style={{background: "#E91E63"}}>
        <div class="navbar-header navbar-brand" style={{color: "white"}}>
          Roetisserie
        </div>
        <div class="navbar-left">
          <ul class="nav navbar-nav">
            {this.renderButtons()}
          </ul>
        </div>
      </div>
    );
  }
}
