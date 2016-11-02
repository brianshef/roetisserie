package llr

import (
    "time"
)

type Item struct {
    id              string
    name            string
    description     string
    clothing_info   ClothingInfo
    inventory_info  InventoryInfo
}

type InventoryInfo struct {
    status                  Status
    //  Cost to seller from source
    cost                    float64
    //  Manufacturer Suggested Retail Price
    msrp                    float64
    //  Actual retail price to customers from seller
    retail_price              float64
    //  Where the item was sourced, eg, "LuLaRoe" or "Jane Doe"
    source                  string
    //  When the item was ordered from the source
    ordered_from_source     time.Time
    //  The tracking number for shipping from source
    source_tracking         string
    //  When the item was received from the source
    received                time.Time
    //  When the item was last placed in IN_STOCK status
    in_stock                time.Time
    //  Whether or not the item may be purchased by customers
    purchaseable            bool
    //  When the item was last placed in FOR_SALE or FOR_SALE_DISCOUNT status
    on_sale                 time.Time
    //  When the item was last placed in PURCHASED status
    customer_purchased      time.Time
    //  When the item was last placed in SHIPPING status
    shipped_to_customer     time.Time
    //  The tracking number for shipping to customer
    customer_tracking       string
    //  When the item was shipped back to seller
    returned                time.Time
    //  The tracking number for returning to seller
    return_tracking         string
}

type ClothingInfo struct {
    style       string
    size        string
    color       string
    pattern     string
    unicorn     bool
    capsule     string
    material    string
}

//
//  STATUS ENUM
//

//  A Status specifies a discrete description of some entity's current status
type Status int
const (
    ORDERED Status = iota   //  Ordered from source
    RECEIVED                //  Received from source
    IN_STOCK                //  Item fully entered into system; not purchaseable
    FOR_SALE                //  Item for sale, can be purchased
    FOR_SALE_DISCOUNT       //  Item for sale at special promotional price
    CLAIMED                 //  Customer has claimed item but not invoiced
    INVOICED                //  Customer has been invoiced for item
    PURCHASED               //  Customer has paid invoice
    PACKAGED                //  Item has been packaged and is ready to ship
    SHIPPING                //  Item shipping to customer, has tracking number
    DELIVERED               //  Item marked as delivered to customer
    RETURNING               //  Item shipping back to seller, has tracking number
    RETURNED                //  Item received back from customer
    DAMAGED                 //  Damage report for item submitted to manufacturer
    DAMAGED_RETURNED        //  Damage item shipped back to manufacturer
    STORAGE                 //  Item is in long-term storage, not for sale
    RESERVED                //  Item reserved for some special purpose
    UNKNOWN                 //  Item's status is unknown or undetermined
    MISSING                 //  Item is missing from seller
    OTHER
)

var statuses = [...]string {
    "ORDERED",
    "RECEIVED",
    "IN STOCK",
    "FOR SALE",
    "FOR SALE DISCOUNTED",
    "CLAIMED",
    "INVOICED",
    "PURCHASED",
    "PACKAGED",
    "SHIPPING",
    "DELIVERED",
    "RETURNING",
    "RETURNED",
    "DAMAGED",
    "DAMAGED RETURNED",
    "IN STORAGE",
    "RESERVED",
    "UNKNOWN",
    "MISSING",
    "OTHER",
}

func (s Status) String() string { return statuses[s] }

//
//  PRIORITY ENUM
//

type Priority int
const (
    NONE Priority = iota
    LOW
    NORMAL
    HIGH
)

var priorities = [...]string {
    "NONE",
    "LOW",
    "NORMAL",
    "HIGH",
}

func (p Priority) String() string { return priorities[p] }
