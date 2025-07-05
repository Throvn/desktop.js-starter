import * as GUI from "GUI"

class MyComponent {
    constructor(props: any) {
        console.log("Hello from my component!", props)
    }

    render() {
        return <vStack>
            <text>Top</text>
            <spacer />
            <hStack>
                <text>Left</text>
                <spacer />
                <text>Center</text>
                <spacer />
                <text>Right</text>
            </hStack>
            <spacer />
            <text>Bottom</text>
        </vStack>
    }
}

GUI.render(<MyComponent />)