import React from "react";
import _ from "lodash";
// import AddRemoveLayout from "./AddRemoveLayout";
// import Gallery from 'react-grid-gallery';
import { CSSGrid, layout, measureItems, makeResponsive } from 'react-stonecutter';

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
      items: [
        {
          url: 'http://remoteaccess1.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess2.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess3.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess5.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess6.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess7.i3s.up.pt',
          scrotPath: './mock.png'
        },
        {
          url: 'http://remoteaccess8.i3s.up.pt',
          scrotPath: './mock.png'
        },
      ]
    };
  }

  renderOverlay(i) {
    return (
      i.tags.map((t) => {
        return (<div
          key={t.value}>
          {t.caption}
        </div>);
      })
    );
  }

  onClickFunc() {
    console.log(this)
    window.open(this.props.item.src, "_blank")
  }

  render() {
    const IMAGES =
      [{
        src: "https://c2.staticflickr.com/9/8817/28973449265_07e3aa5d2e_b.jpg",
        thumbnail: "https://c2.staticflickr.com/9/8817/28973449265_07e3aa5d2e_n.jpg",
        thumbnailWidth: 320,
        thumbnailHeight: 212,
        caption: "After Rain (Jeshu John - designerspics.com)"
      },
      {
        src: "https://c2.staticflickr.com/9/8356/28897120681_3b2c0f43e0_b.jpg",
        thumbnail: "https://c2.staticflickr.com/9/8356/28897120681_3b2c0f43e0_n.jpg",
        thumbnailWidth: 320,
        thumbnailHeight: 212,
        tags: [{ value: "i3s.up.pt", title: "Ocean" }, { value: "People", title: "People" }],
        caption: "Boats (Jeshu John - designerspics.com)"
      },

      {
        src: "https://c4.staticflickr.com/9/8887/28897124891_98c4fdd82b_b.jpg",
        thumbnail: "https://c4.staticflickr.com/9/8887/28897124891_98c4fdd82b_n.jpg",
        thumbnailWidth: 320,
        thumbnailHeight: 212
      }]
    const Grid = makeResponsive(measureItems(CSSGrid), {
      maxWidth: 1920,
      minPadding: 100
    });
    return (
      <div>
        {/* <AddRemoveLayout items={this.state.items} /> */}
        {/* <Gallery images={IMAGES}
                         onClickThumbnail={this.onClickFunc}
                /> */}
        <Grid
          component="ul"
          columns={5}
          columnWidth={150}
          gutterWidth={5}
          gutterHeight={5}
          layout={layout.pinterest}
          duration={800}
          easing="ease-out"
        >
          <li key="A" itemHeight={150}>A</li>
          <li key="B" itemHeight={120}>B</li>
          <li key="C" itemHeight={170}>C</li>
        </Grid>
      </div>
    );
  }
}