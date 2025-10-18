/**
 * These are the types which are available in the Desktop.js runtime.
 * Check it out: https://github.com/Throvn/desktop.js
 * Do not modify this file it could be subject to change.
 */

export { };

type HEX = `#${string}`;
type NamedColors = "transparent" | "aliceblue" | "antiquewhite" | "aqua" | "aquamarine" | "azure" | "beige" | "bisque" | "black" | "blanchedalmond" | "blue" | "blueviolet" | "brown" | "burlywood" | "cadetblue" | "chartreuse" | "chocolate" | "coral" | "cornflowerblue" | "cornsilk" | "crimson" | "cyan" | "darkblue" | "darkcyan" | "darkgoldenrod" | "darkgray" | "darkgreen" | "darkgrey" | "darkkhaki" | "darkmagenta" | "darkolivegreen" | "darkorange" | "darkorchid" | "darkred" | "darksalmon" | "darkseagreen" | "darkslateblue" | "darkslategray" | "darkslategrey" | "darkturquoise" | "darkviolet" | "deeppink" | "deepskyblue" | "dimgray" | "dimgrey" | "dodgerblue" | "firebrick" | "floralwhite" | "forestgreen" | "fuchsia" | "gainsboro" | "ghostwhite" | "gold" | "goldenrod" | "gray" | "green" | "greenyellow" | "grey" | "honeydew" | "hotpink" | "indianred" | "indigo" | "ivory" | "khaki" | "lavender" | "lavenderblush" | "lawngreen" | "lemonchiffon" | "lightblue" | "lightcoral" | "lightcyan" | "lightgoldenrodyellow" | "lightgray" | "lightgreen" | "lightgrey" | "lightpink" | "lightsalmon" | "lightseagreen" | "lightskyblue" | "lightslategray" | "lightslategrey" | "lightsteelblue" | "lightyellow" | "lime" | "limegreen" | "linen" | "magenta" | "maroon" | "mediumaquamarine" | "mediumblue" | "mediumorchid" | "mediumpurple" | "mediumseagreen" | "mediumslateblue" | "mediumspringgreen" | "mediumturquoise" | "mediumvioletred" | "midnightblue" | "mintcream" | "mistyrose" | "moccasin" | "navajowhite" | "navy" | "oldlace" | "olive" | "olivedrab" | "orange" | "orangered" | "orchid" | "palegoldenrod" | "palegreen" | "paleturquoise" | "palevioletred" | "papayawhip" | "peachpuff" | "peru" | "pink" | "plum" | "powderblue" | "purple" | "red" | "rosybrown" | "royalblue" | "saddlebrown" | "salmon" | "sandybrown" | "seagreen" | "seashell" | "sienna" | "silver" | "skyblue" | "slateblue" | "slategray" | "slategrey" | "snow" | "springgreen" | "steelblue" | "tan" | "teal" | "thistle" | "tomato" | "turquoise" | "violet" | "wheat" | "white" | "whitesmoke" | "yellow" | "yellowgreen";
type Padding = number | {
    /**
     * Applies padding to the **left** and **right**.
     * Individual sides take precedent.
     */
    horizontal?: number;
    /**
     * Applies padding to the **top** and **bottom**.
     * Individual sides take precedent.
     */
    vertical?: number;
    left?: number;
    right?: number;
    top?: number;
    bottom?: number;
};

type MouseEvent = {
    /**
     * The `x` coordinate from the left window edge to the mouse.
     * Always positive.
     */
    layerX: number;
    /**
     * The `y` coordinate from the top window edge to the mouse.
     * Always positive.
     */
    layerY: number;
    /**
     * Indicates if the left or right ALT key was pressed during the mouse event.
     * 
     * On macOS ALT is the Option (`⌥`) key.
     */
    altKey: boolean;
    /**
     * Indicates if the left or right CTRL (`⌃`) key was pressed during the mouse event.
     */
    ctrlKey: boolean;
    /**
     * Indicates if the SHIFT (`⇧`) key was pressed during the mouse event.
     */
    shiftKey: boolean;
}

type MouseEventWithButton = MouseEvent & {
    /**
     * The index of the button which was pressed or released.
     */
    button: MouseBtn;
};

type MouseDownEvent = MouseEventWithButton;
type MouseUpEvent = MouseEventWithButton;

/**
 * Determines how rounded the corners of an element (e.g. stack) are. 
 * 
 * Union radii (like top or bottom) are evaluated first, 
 * which makes it possible to overwrite part of the values 
 * through more fine grained properties like `topLeft`.
 */
interface BorderRadius {
    top?: number;
    bottom?: number;
    left?: number;
    right?: number;
    topLeft?: number;
    topRight?: number;
    bottomLeft?: number;
    bottomRight?: number;
}

type TextProps = {
    id?: string;
    $borderRadius?: number | BorderRadius;
    $backgroundColor?: HEX | NamedColors;
    /**
     * Applied to the text itself.
     * Supported color types are: HTML Colors or hex colors with or without an alpha value.
     */
    $color?: HEX | NamedColors;
    /**
     * Font size is generally thought of as `x pixels tall`. Default is 12 (pixels tall).
     */
    $fontSize?: number;
    /**
     * The file name of the font file which is located under `assets/fonts`.
     * 
     * E.g.
     * assets/fonts/Roboto-Regular.ttf
     * 
     * ```js
     * <text $fontFace="Roboto-Regular.ttf">I have a different font face!</text>
     * ```
     */
    $fontFace?: string;
    /**
     * Results in horizontal whitespace between the individual characters.
     */
    $letterSpacing?: number;
    /**
     * Determines how tall the text should be. Default is 12 (just like the font size).
     */
    $lineHeight?: number;
    $padding?: Padding;
    onMouseDown?: (event: MouseDownEvent) => void;
    onMouseOver?: (event: MouseEvent) => void;
    onMouseUp?: (event: MouseUpEvent) => void;
}

type StackProps = {
    id?: string;
    $borderRadius?: number | BorderRadius;
    $backgroundColor?: HEX | NamedColors;
    /**
     * Adds `x px` of spacing between the child elements of a stack.
     */
    $gap?: number;
    $padding?: Padding;
    onMouseDown?: (event: MouseDownEvent) => void;
    onMouseOver?: (event: MouseEvent) => void;
    onMouseUp?: (event: MouseUpEvent) => void;
}

declare global {
    enum MouseBtn {
        left = 0,
        right = 1,
        middle = 2,
    }

    namespace JSX {
        interface IntrinsicElements {
            /**
            A container component which applies all of its props to each child individually. E.g. useful for styling list items.

            - If the same property was already defined on the child, the property is **NOT** applied. As props closer to the component always take precedent.
            - If a child component doesn't specify the style, it is ignored by the child component.
            - It doesn't contribute an element to the view hierarchy. (Similar to [React.js Fragment](https://react.dev/reference/react/Fragment))

            ##### Example

            When you have the following code:

            ```html
            <vStack>
            <group $backgroundColor="black">
                <text>First Text</text>
                <text>Second Text</text>
            </group>
            </vStack>
            ```

            You can think of it as:

            ```html
            <vStack>
            <text $backgroundColor="black">First Text</text>
            <text $backgroundColor="black">Second Text</text>
            </vStack>
            ```
             */
            "group": Object,
            /**
             * A component to style text. All of its children **NEED** to be strings, otherwise, _undefined_ will be printed on the screen.
             */
            "text": TextProps;
            /**
             * This container component orders its child components **vertically** (from top to bottom).
             *
             * Each child is centered automatically. To change the position of children use [`<spacer />`s](#spacer-).
             */
            "vStack": StackProps,
            /**
             * This container component orders its child components **horizontally** (from left to right).
             *
             * Each child is centered automatically. To change the position of children use [`<spacer />`s](#spacer-).
             */
            "hStack": StackProps,

            /**
             * A component which takes up all available space inside of its parent component.
             * Use it to position children of `<hStack>` and `<vStack>`.
             * 
             * This component doesn't take any props.
             */
            "spacer": {},

            /**
             * A component to display an image. It's children are displayed only if the the supplied `data` prop is not a valid `Blob` object, functioning as a placeholder.
             *
             * If the `$width` and or `$height` attributes are set before the image loaded, the placeholder will receive the same dimensions.
             * 
             * #### Supported image types
             * - `image/bmp`
             * - `image/gif` Only shows the first frame. Not moving images!
             * - `image/jpeg`
             * - `image/png`
             * - `image/psd`
             */
            "img": {
                $width?: number,
                $height?: number,
                data: any, // TODO: Make this the Blob type!
            }
        }
    }

    /**
     * Standard console.
     */
    const console: {
        /**
         * Outputs a message to the console.
         * Can handle multiple arguments of any type.
         */
        log(...data: any[]): void;
    }
}

