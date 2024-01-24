export interface Card {
    id: string
    name: string
    description: string
    typeid: string
    hp: number
    attack: number
    defense: number
    price: number
    ownerid: string
    islocked: boolean
    image: string
    typename?: string
}