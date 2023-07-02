import TextField from "../components/TextField";
import TaskDisplayArea from "../components/TaskDisplayArea";
import Button from "../components/Button";

const Home = () => {
    return (
        <div className="h-100 w-full flex items-center justify-center bg-teal-lightest font-sans">
            <div className="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
                <div className="mb-4">
                    <h1 className="text-grey-darkest">Todo List</h1>
                    <div className="flex mt-4">
                        <TextField />
                        <Button label={'Add'} />
                    </div>
                </div>
                <TaskDisplayArea />
            </div>
        </div>
    )
}

export default Home
