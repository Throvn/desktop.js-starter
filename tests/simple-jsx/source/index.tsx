import * as another from "./another"

const nested = <text>This is a test</text>
let fragment = <group> <vStack> <spacer /> </vStack>Test</group>

console.log(nested, another.b, fragment)
class MyComponent {
    constructor(props: any) {

    }

    render() {
        return <text>This is my component!</text>
    }
}

GUI.render(<MyComponent />)