import Button from "./Button";
import {useEffect, useState} from "react";
import apis, {GetTodoResponse} from "../lib/apis";

const TaskDisplayArea = () => {
    const [todos, setTodos] = useState<GetTodoResponse>()
    useEffect(() => {
        const fetchTodos = async () => {
            setTodos(await apis.getTodo())
        }
        fetchTodos()
    }, [setTodos])

    return (
        <div>
            <div className="flex mb-4 items-center">
                { todos?.data.map((title, key) =>
                    <p key={ key }>{ title }</p>
                ) }
                {/*<p className="w-full text-grey-darkest">Add another component to Tailwind Components</p>*/}
                {/*<Button label={'Done'} customClass={'mr-5'} />*/}
                {/*<Button label={'Remove'} />*/}
            </div>
        </div>
    )
}

export default TaskDisplayArea
