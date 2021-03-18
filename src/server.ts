import express, { Application } from "express";

const PORT: any | number = process.env.PORT || 8088;   
const app: Application = express();

app.get("/ping", async (req, res) => {
    console.log("ping");
    res.send("pong");
});

app.listen(PORT, () => {
    console.log("Server is running on port: ", PORT);
});