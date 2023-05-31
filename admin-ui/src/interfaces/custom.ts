export interface ChatMessage {
  id: string;
  message?: string;
  parent?: string | null;
  children?: Array<string>;
  typing?: boolean;
  content?: string;
};
