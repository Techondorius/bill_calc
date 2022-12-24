import { MouseEventHandler, ReactNode } from 'react'

type Props = {
  children?: ReactNode
  className?: string
  type: string
  onClick?: MouseEventHandler<HTMLButtonElement>
}

export const MyButton: React.FC<Props> = ({ children, className, type, onClick }) => {
  const classes =
    type != 'main'
      ? className +
        ' shadow bg-white border-2 border-blue-400 hover:bg-blue-200 focus:shadow-outline focus:outline-none text-blue-400 py-2 px-4 rounded'
      : className +
        ' shadow border-2 border-blue-500 bg-blue-500 hover:bg-blue-400 focus:shadow-outline focus:outline-none text-white py-2 px-4 rounded'

  return (
    <button className={classes} type='button' onClick={onClick}>
      {children}
    </button>
  )
}
