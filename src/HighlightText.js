import React, { Component } from 'react';

class HighlightText extends Component {
  render() {
    const highlights = this.props.text.split('>>');
    const first = highlights.shift();
    return (
      <span>
        {first}
        {highlights.map(s => s.split('<<', 2)).map((s, i) =>
          <span key={i}>
            <span style={{color: 'red'}}>{s[0]}</span>
            {s[1]}
          </span>
        )}
      </span>
    )
  }
}

export default HighlightText;
