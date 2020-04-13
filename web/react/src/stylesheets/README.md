Style architecture based on [ITCSS](https://www.creativebloq.com/web-design/manage-large-css-projects-itcss-101517528)
and [RSCSS](https://rscss.io/).

- `settings` - Basic architecture configurations, global variables defining colors, spacing, and so on
- `tools` - Mixins and functions used in the project
- `generic` - Most generic properties with the least specificity. Resets, box-sizing, and so on
- `elements` - Basic styles for HTML elements (`h1`, `button`, ...). This is the last layer where selectors will be applied directly in tags.
- `objects` - Common interface components, like lists, panels, buttons, and so on
- `components` - Application-specific reusable components, adhering to RSCSS rules. Product listing, specific cards, so on
- `trumps` - Most specific layer, overrides everything above. Contains utility and helper classes, hacks and overrides.
