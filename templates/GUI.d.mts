declare module "GUI" {
    interface InitialRenderOptions {
        /**
         * The title of your window. Defaults to 'Desktop.js'.
         */
        title?: string
        /**
         * The initial width of the window in pixels.
         */
        width?: number
        /**
         * The initial height of the window in pixels.
         */
        height?: number
        /**
         * Whether the window is resizable by the user.
         */
        resizable?: boolean
    }
    /**
     * Whatever gets passed in here is rendered to the screen.
     * Only call this function at one point in time in your app.
     * @param root Entry point of your render tree.
     * @param options Initial window options object.
     */
    function render(root: JSX.IntrinsicElements, options?: InitialRenderOptions): void;
}