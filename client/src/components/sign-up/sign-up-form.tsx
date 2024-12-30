import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

export const SignUpForm = () => {
  return (
    <form className="grid grid-cols-1 gap-4">
      <div>
        <Label>Name</Label>
        <Input type="text" placeholder="Enter your name" />
      </div>
      <div>
        <Label>Email</Label>
        <Input type="text" placeholder="Enter your email" />
      </div>
      <div>
        <Label>Password</Label>
        <Input type="password" placeholder="Enter your password" />
      </div>
      <div>
        <Label>Confirm Password</Label>
        <Input type="password" placeholder="Confirm your password" />
      </div>
      <Button type="submit" className="w-full mt-4">
        Continue
      </Button>
    </form>
  );
};
