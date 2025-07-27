import { useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";

const InviteRedirect = () => {
  const { token } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    if (token) {
      localStorage.setItem("invite_token", token);
    }

    // Optional: delay to show a "redirecting..." message
    navigate("/home");
  }, [token, navigate]);

  return null; // or a loading spinner if you want
};

export default InviteRedirect;
