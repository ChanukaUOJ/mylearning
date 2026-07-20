import { useEffect, useState } from "react";

export default function UseEffect(){
    const [posts, setPosts] = useState([])

    useEffect(() => {
        // use effects shows this console log in the browser console twice because of some internal react behaviour(React strict mode).
        // This wont run twice in the production environment.
        console.log("Page Rendered")

        // api call
        fetch("https://jsonplaceholder.typicode.com/posts").then((response) => response.json()).then((json) => setPosts(json));

        // this dependency array is to re run the use effect when a state change of a given state
        // ex: below array --> [counter]
    },[])

    return (
        <div>
            <h1>Posts</h1>
            <ul>
                {posts && posts.map((post) => (<li key={post.id}>{post.title}</li>))}
            </ul>
        </div>
    )
}