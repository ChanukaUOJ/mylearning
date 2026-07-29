import { forwardRef, useImperativeHandle, useRef } from "react";

export default function UseImperativeHandle() {
    const inputRef = useRef();

  return (
    <div className="gap-4 flex">
    <CustomInput ref={inputRef} />
      <button onClick={() => inputRef.current.focusInput()}>Focus Input</button>
      <button onClick={() => inputRef.current.clearInput()}>Clear Input</button>
    </div>
  );
}

const CustomInput = forwardRef((props, ref) => {
    const inputRef = useRef();

    useImperativeHandle(ref, () => ({
        focusInput: () => {
            inputRef.current.focus()
        },
        clearInput: () => {
            inputRef.current.value = ""
        },
    }))

    return <input ref={inputRef} type="text" placeholder="Type something..." />
})
