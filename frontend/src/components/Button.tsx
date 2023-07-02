import {ButtonHTMLAttributes, DetailedHTMLProps} from "react"

interface Props {
    label: string
    children?: Element
    customClass?: string
}

type ButtonProps = DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>

const Button = ({
                    label,
                    children,
                    customClass,
                    ...buttonProps
                }: Props & ButtonProps) => {
    return (
        <button
            className={'flex-no-shrink p-2 border-2 rounded text-teal border-teal hover:text-white hover:bg-teal ' +
                customClass
            } {...buttonProps}>
            {children ? children : label}
        </button>
    )
}

export default Button
