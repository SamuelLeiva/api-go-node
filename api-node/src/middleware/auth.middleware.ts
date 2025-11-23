import { NextFunction, Request, Response } from "express";
import jwt from "jsonwebtoken";

export function authRequired(req: Request, res: Response, next: NextFunction){
    const secret = process.env.JWT_SECRET; //|| "superSecretKey518485151H8t+ds*&^%$#@!";
    console.log("Using JWT secret:", secret);
    if(!secret){
        return res.status(500).json({message: "JWT secret not configured"});
    }

    const header = req.headers['authorization'];
    if(!header){
        return res.status(401).json({message: "Authorization header missing"});
    }

    const token = header.replace('Bearer ', '');

    try {
        jwt.verify(token, secret);
        return next();
    } catch (err) {
        return res.status(401).json({message: "Invalid or expired token"});
    }
}