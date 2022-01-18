import { decode, encode } from './xmorse.min';
import config from "./config.json";

const token = config["token"];
const bot_name = config["bot_name"];
const master_id = config["master_id"];

addEventListener("fetch", (event) => {
  event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
  if (request.method === "POST") {
    let data = await request.json();
    console.log(data.message);
    let message_id = data.message.message_id;
    let chat_id = data.message.chat.id;
    let text = data.message.text || "";
    let texts = text.split(" ");
    if (text[0] === "/") {
      texts[0] = texts[0].replace("/", "").replace(bot_name, "");
      switch (texts[0]) {
        case "help":
        case "start":
          await tg(token, "sendMessage", {
            chat_id: chat_id,
            text: ";)",
          });
          break;
        case "decode":
          let decodeMessage = decode(texts[1]);
          await tg(token, "sendMessage", {
            chat_id: chat_id,
            reply_to_message_id: message_id,
            text: `${decodeMessage}`,
          });
          break;
        case "encode":
          let encodeMessage = encode(texts[1]);
          await tg(token, "sendMessage", {
            chat_id: chat_id,
            reply_to_message_id: message_id,
            text: `${encodeMessage}`,
          });
          break;
      }
    }
  }

  return new Response("ok", { status: 200 });
}

async function tg(token, type, data) {
  let response = await (
    await fetch(`https://api.telegram.org/bot${token}/${type}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
  ).json();
  if (!response.ok) {
    await tg(token, "sendMessage", {
      chat_id: master_id,
      text: JSON.stringify(response),
    });
  }
}
