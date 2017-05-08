import React, { Component } from 'react';
import './ResultTable.css';

function ValueCell({ data }) {
  if(data === null) {
    return (
      <span>(null)</span>
    );
  }

  const maxLines = 4;
  const maxLength = 10;

  let dataLines = `${data}`.split(/\r?\n/);

  if(dataLines.length > maxLines) {
    dataLines = dataLines.slice(0, maxLines);
    dataLines.push('...');
  }

  dataLines = dataLines.map(l => {
    if(l.length > maxLength) {
      l = l.substring(0, maxLength) + "...";
    }
    return l;
  });

  const preview = dataLines.reduce((prev, curr) => [prev, <br key={curr} />, curr])

  return (
    <span title={data}>{preview}</span>
  );
}

class ResultTable extends Component {
  copyTable() {
    if (typeof document !== 'undefined'
      && typeof window !== 'undefined'
      && document.createRange
      && document.execCommand
      && window.getSelection) {

      const sel = window.getSelection();
      sel.removeAllRanges();

      const range = document.createRange();
      range.setStartBefore(this.tHead);
      range.setEndAfter(this.tBody);
      sel.addRange(range);
      document.execCommand('copy');
    }
  }

  render() {
    const { caption, data } = this.props;

    if (!data || data.length === 0) {
      return (
        <table className="table">
          { caption ? <caption>{caption}</caption> : ""}
          <tbody>
            <tr><td>No data returned in result.</td></tr>
          </tbody>
        </table>
      );
    }

    const cols = Object.keys(data[0])

    return (
      <table ref={r => { this.table = r; }} className="table table-striped table-condensed small result-table">
        <caption>
          { caption }
          <br />
          <small>{ data.length } rows</small>

          <button className="btn btn-xs" type="button" onClick={() => this.copyTable()}>Copy</button>
        </caption>
        <thead ref={r => { this.tHead = r; }}>
          <tr>
            {cols.map(c =>
              <th key={c}>{c}</th>
            )}
          </tr>
        </thead>
        <tbody ref={r => { this.tBody = r; }}>
          {data.map((r, i) =>
            <tr key={i}>
              {cols.map(c =>
                <td key={c}>
                  <ValueCell data={r[c]} />
                </td>
              )}
            </tr>
          )}
        </tbody>
      </table>
    )
  }
}
export default ResultTable;
