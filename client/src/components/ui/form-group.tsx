import { Label } from "./label";

interface Props {
  children: React.ReactNode;
  label: string;
  error?: string;
}

export const FormGroup = ({ children, label, error }: Props) => {
  return (
    <div>
      <Label>{label}</Label>
      {children}
      {error && <span className="text-xs text-red-500 ml-1">{error}</span>}
    </div>
  );
};
