import { CardType } from "./cardtype.interface"


export interface TypeResponse {
    types: CardType[]
    errors: string[]
}