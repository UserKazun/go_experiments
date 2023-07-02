import Button from "./Button";

const TaskDisplayArea = () => {
    return (
        <div>
            <div className="flex mb-4 items-center">
                <p className="w-full text-grey-darkest">Add another component to Tailwind Components</p>
                <Button label={'Done'} customClass={'mr-5'} />
                <Button label={'Remove'} />
            </div>
            <div className="flex mb-4 items-center">
                <p className="w-full line-through text-green">Submit Todo App Component to Tailwind
                    Components</p>
                <Button label={'Not Done'} customClass={'mr-5'} />
                <Button label={'Remove'} />
            </div>
        </div>
    )
}

export default TaskDisplayArea
