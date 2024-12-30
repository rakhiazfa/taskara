import { Button } from "../ui/button";
import { Checkbox } from "../ui/checkbox";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

export const SignInForm = () => {
  return (
    <form className="grid grid-cols-1 gap-4">
      <div>
        <Label>Email</Label>
        <Input type="text" placeholder="Enter your email" />
      </div>
      <div>
        <Label>Password</Label>
        <Input type="password" placeholder="Enter your password" />
      </div>
      <div className="flex items-center gap-2">
        <Checkbox />
        <Label>Remember me</Label>
      </div>
      <Button type="submit" className="w-full">
        Continue
      </Button>
    </form>
  );
};
