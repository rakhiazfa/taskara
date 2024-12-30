import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

export const ResetPasswordForm = () => {
  return (
    <form className="grid grid-cols-1 gap-4">
      <div>
        <Label>New Password</Label>
        <Input type="password" placeholder="Enter new password" />
      </div>
      <div>
        <Label>Confirm New Password</Label>
        <Input type="password" placeholder="Confirm new password" />
      </div>
      <Button type="submit" className="w-full mt-4">
        Continue
      </Button>
    </form>
  );
};
