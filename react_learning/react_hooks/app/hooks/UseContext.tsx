import React, { useState, createContext, useContext } from "react";

interface ContextProps {
  toggle: boolean;
  setToggle: React.Dispatch<React.SetStateAction<boolean>>;
}

const GlobalToggleContext = createContext<ContextProps>({
  toggle: false,
  setToggle: () => {},
});

export default function UseContext() {
  const [toggle, setToggle] = useState(false);

  return (
    <GlobalToggleContext.Provider value={{ toggle, setToggle }}>
      <div className="text-center mt-20">
        <p className="text-2xl">Parent Componet:</p>
        <ToggleButton />
        <DisplayBox />
      </div>
    </GlobalToggleContext.Provider>
  );
}

function ToggleButton() {
  const { toggle, setToggle } = useContext(GlobalToggleContext);

  return (
    <button
      onClick={() => setToggle((prev) => !prev)}
      className="border p-4 rounded-full my-5 cursor-pointer"
    >
      Change State to: {toggle ? "true" : "false"}
    </button>
  );
}

function DisplayBox() {
  const { toggle } = useContext(GlobalToggleContext);

  return <p className="text-red-500">Button Value is : {toggle ? "true" : "false"}</p>;
}
