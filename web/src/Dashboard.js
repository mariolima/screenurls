import React from "react";
import _ from "lodash";
import AddRemoveLayout from "./AddRemoveLayout";

/**
 * This layout demonstrates how to use a grid with a dynamic number of elements.
 */
export default class Dashboard extends React.PureComponent {
  static defaultProps = {
    className: "layout",
    cols: { lg: 12, md: 10, sm: 6, xs: 4, xxs: 2 },
    rowHeight: 100
  };

  constructor(props) {
    super(props);

    this.state = {
      items: [{
        url: 'http://remoteaccess.i3s.up.pt',
        scrotPath: './mock.png'
      }]
    };
  }

  render() {
    return (
      <div>
        <AddRemoveLayout items={this.state.items}/>
      </div>
    );
  }
}

// import("./test-hook.js").then(fn => fn.default(AddRemoveLayout));