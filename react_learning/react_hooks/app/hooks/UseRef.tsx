import { useEffect, useRef, useState } from "react"

export default function UseRef(){
    const [count, setCount] = useState(0)
    const inputRef = useRef<HTMLInputElement>(null);

    useEffect(()=>{
        console.log('initial use effect')
        if(inputRef.current){
            console.log(inputRef.current.value)
        }
    },[])

    // use effect to increment count
    useEffect(()=>{
        console.log('count variable value : ',{count})
    },[count])

    const handleClick = () => {
        if(inputRef.current){
            // inputRef.current.value = "Chamith"
            inputRef.current.focus();
            console.log('handle click value : ',inputRef.current.value)
        }
    }

    const handleCountClick = () => {
        setCount((prev) => prev + 1)
    }
    
    return (
        <>
            <input id="name" type="text" ref={inputRef} autoComplete="true"/>
            <button onClick={handleClick}>Click me</button>

            <p>{count}</p>
            <button onClick={handleCountClick}>Increment</button>
        </>
    )
}