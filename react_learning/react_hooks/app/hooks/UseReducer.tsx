import { useReducer } from "react";

// use reducer is an alternative for useState which helps you to handle more complex states
// here is the boilerplater useState use case
// const buttonStyle = "px-3 py-1 border gap-2 cursor-pointer"

// export default function UseState() {
//   const [count, setCount] = useState(0);

//   return (
//     <div className="flex justify-center h-screen items-center">
//       <div className="gap-2">
//         <p>Count : {count}</p>
//         <div>
//           <button className={buttonStyle} onClick={() => setCount(count + 1)}>+</button>
//           <button className={buttonStyle} onClick={() => setCount(count - 1)}>-</button>
//           <button className={buttonStyle} onClick={() => setCount(count * 2)}>*</button>
//         </div>
//       </div>
//     </div>
//   );
// }

// use reducter alternative
const buttonStyle = "px-3 py-1 border gap-2 cursor-pointer"

interface State{
    count: number;
}

type Action = {type: "increment"} | {type: "decrement"} | {type: "multiply"}

const reducer = (state: State, action: Action): State => {
    switch(action.type){
        case "increment":
            return {count: state.count + 1};
        case "decrement":
            return {count: state.count - 1};
        case "multiply":
            return {count: state.count * 2};
        default:
            return {count: 0}
    }
}

export default function UseReducer() {
  const [state, dispatch] = useReducer(reducer, {count: 0})

  return (
    <div className="flex justify-center h-screen items-center">
      <div className="gap-2">
        <p>Count : {state.count}</p>
        <div>
          <button className={buttonStyle} onClick={() => dispatch({type: "increment"})}>+</button>
          <button className={buttonStyle} onClick={() => dispatch({type: "decrement"})}>-</button>
          <button className={buttonStyle} onClick={() => dispatch({type: "multiply"})}>*</button>
        </div>
      </div>
    </div>
  );
}
