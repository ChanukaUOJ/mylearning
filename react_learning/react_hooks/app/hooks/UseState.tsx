import {useState} from "react";

export default function UseState(){

    // to keep the internal state of the component

    const [count, setCount] = useState(0);

    // handle increment
    const handleIncrement = () => {
        setCount((prev) => prev + 1)
    }

    // handle decrement
    const handleDecrement = () => {
        setCount((prev) => prev + 1)
    }

    const buttonStyle = "cursor-pointer px-2"

    return (
        <div>
            <p>Count : {count}</p>
            <button className={buttonStyle} onClick={handleIncrement}>Increase Counter</button>
            <button className={buttonStyle} onClick={handleDecrement}>Decrease Counter</button>
        </div>
    )
}