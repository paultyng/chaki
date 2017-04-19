import React, { Component } from 'react';

import './HighlightText.css';

class HighlightText extends Component {
  render() {
    const highlights = this.props.text.split('>>');
    const first = highlights.shift();
    return (
      <span className="highlight-text">
        {first}
        {highlights.map(s => s.split('<<', 2)).map((s, i) =>
          <span key={i}>
            <mark>{s[0]}</mark>
            {s[1]}
          </span>
        )}
      </span>
    )
  }
}

export default HighlightText;
