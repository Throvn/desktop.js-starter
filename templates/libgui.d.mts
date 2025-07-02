declare module "libgui" {
    /**
     * Whatever gets passed in here is rendered to the screen.
     * Only call this function at one point in time in your app.
     * @param root Entry point of your render tree
     */
    function render(root: JSX.IntrinsicElements): void;
}