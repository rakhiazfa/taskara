interface Props {
  className?: string;
}

export const AppLogo = ({ className }: Props) => (
  <div className={className}>
    <span className="text-xl font-black">
      {process.env.NEXT_PUBLIC_APP_NAME}
    </span>
  </div>
);
