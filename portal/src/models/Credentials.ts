export type credentials = {
  username: string;
  password: string;
};

export function buildCredentials(raw: any): credentials {
  return {
    username: raw.username,
    password: raw.password
  };
}

