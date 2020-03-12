import React from "react";
import { WidthProvider, Responsive } from "react-grid-layout";
import _ from "lodash";
const ResponsiveReactGridLayout = WidthProvider(Responsive);

/**
 * This layout demonstrates how to use a grid with a dynamic number of elements.
 */
export default class AddRemoveLayout extends React.PureComponent {
  static defaultProps = {
    className: "layout",
    cols: { lg: 12, md: 10, sm: 6, xs: 4, xxs: 2 },
    onLayoutChange: function() {},
    rowHeight: 100,
  };

  constructor(props) {
    super(props);

    this.state = {
      items: [],
    };

    this.onBreakpointChange = this.onBreakpointChange.bind(this);
  }

  createElement(el,k) {
    el.x=(k * 2);
    el.y=Infinity; // puts it at the bottom
    el.w=2;
    el.h=2;
    const infoStyle = {
      position: "absolute",
      left: "2px",
      bottom: 0,
      cursor: "pointer",
      color: `black`
    };
    const i = k
    const mstyle = {
      backgroundImage: `url(${el.scrotPath})`,
      backgroundPosition: 'center',
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat'
    };
    return (
      <div key={i} data-grid={el} style={mstyle} >
        <span className="text">{i}</span>
        <span
          className="remove"
          style={infoStyle}
          onClick={this.onRemoveItem.bind(this, i)}
        >
          {el.url}
        </span>
      </div>
    );
  }

  // We're using the cols coming back from this to calculate where to add new items.
  onBreakpointChange(breakpoint, cols) {
    this.setState({
      breakpoint: breakpoint,
      cols: cols
    });
  }

  onLayoutChange(layout) {
    this.props.onLayoutChange(layout);
    this.setState({ layout: layout });
  }

  onRemoveItem(i) {
    console.log("removing", i);
    this.setState({ items: _.reject(this.state.items, { i: i }) });
  }

  onAddItem() {
    console.log(this.state)
  }

  render() {
    this.state.items=this.props.items
    return (
      <div>
        <button onClick={this.onAddItem}>Add Item</button>
        <ResponsiveReactGridLayout
          onLayoutChange={this.onLayoutChange}
          onBreakpointChange={this.onBreakpointChange}
        >
          {_.map(this.props.items, (el,k) => this.createElement(el,k))}
        </ResponsiveReactGridLayout>
      </div>
    );
  }
}

// import("./test-hook.js").then(fn => fn.default(AddRemoveLayout));