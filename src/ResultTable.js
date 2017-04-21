import React, { Component } from 'react';

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
      <table className="table table-striped table-condensed small">
        <caption>
          { caption }
          <br />
          <small>{ data.length } rows</small>
        </caption>

        <thead>
          <tr>
            <th>#</th>
            {cols.map(c =>
              <th key={c}>{c}</th>
            )}
          </tr>
        </thead>
        <tbody>
          {data.map((r, i) =>
            <tr key={i}>
              <td>{i + 1}</td>
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
