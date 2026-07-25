import UseState from "~/hooks/UseState";
import type { Route } from "./+types/home";
import UseEffect from "~/hooks/UseEffect";
import UseContext from "~/hooks/UseContext";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  return (
    // use state
    // <UseState/>

    // use effect
    // <UseEffect/>

    // use context
    <UseContext/>
  );
}
