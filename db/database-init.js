// --- Init User collection ---
var tokensDefinition = {
    bsonType: "array",
    items: {
        bsonType: "object",
        required: [ "loginTime", "authDevice", "jwtToken" ],
        properties: {
            loginTime: {
                bsonType: "date"
            },
            logoutTime: {
                bsonType: "date"
            },
            authDevice: {
                bsonType: "string"
            },
            jwtToken: {
                bsonType: "string"
            }
        }
    }
};

var rolesDefinition = {
    bsonType: "array",
    items: {
        bsonType: "object",
        required: [ "name", "isRemovable" ],
        properties: {
            name: {
                bsonType: "string"
            },
            code: {
                bsonType: "long"
            },
            isRemovable: {
                bsonType: "boolean"
            }
        }
    },
    minItems: 1
};

var userCollectionFieldValidator = {
    $jsonSchema: {
        bsonType: "object",
        required: [ "nickname", "email", "password" ],
        properties: {
            nickname: {
                bsonType: "string"
            },
            email: {
                bsonType: "string",
                pattern: "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+(?:[A-Z]{2}|com|org|net|gov|mil|biz|info|mobi|name|aero|jobs|museum|by|ru)\b"
            },
            password: {
                bsonType: "string"
            },
            firstName: {
                bsonType: "string"
            },
            lastName: {
                bsonType: "string"
            },
            gender: {
                enum: ["M", "F"]
            },
            phoneNumber: {
                bsonType: "string"
            },
            createdTime: {
                bsonType: "date"
            },
            tokens: tokensDefinition,
            roles: rolesDefinition
        }
    }
};

// --- Create collection ---

db.createCollection("users", {
    validator: userCollectionFieldValidator,
    validationAction: "error",
    validationLevel: "strict",
});

// --- Create indexes ---

db.users.createIndex({ "nickname": 1 }, { unique: true });
db.users.createIndex({ "email": 1 }, { unique: true });

// -- Init data ---

var adminRole = {
    name: "Admin",
    code: 0,
    isRemovable: false
};

var userRole = {
    name: "User",
    code: 1,
    isRemovable: false
};

var companyOwnerRole = {
    name: "Company Owner",
    code: 2,
    isRemovable: true
};

var companyManagerRole = {
    name: "Company Manager",
    code: 3,
    isRemovable: true
};

var admin = {
    nickname: "root",
    email: "admin@email.com",
    password: "$2a$04$u3MXPUix1X8Lg8b8AK4lZOIRCDLZmj/cI0UlHA4Ri2LSBMSBEvpAu",
    gender: "M",
    roles: [ adminRole ]
};

db.users.insert(admin);

// TODO: For Role entity need to create isDeletable field that will explain that it "default" roles for user, which we can't remove from user.

// Do we really need this?
// --- Init Roles collection ---
// var roleCollectionFieldValidator = {
//     $jsonSchema: {
//         bsonType: "object",
//         required: [ "name" ],
//         properties: {
//             name: {
//                 bsonType: "string"
//             }
//         }
//     }
// }

// db.createCollection("roles", {
//     validator: roleCollectionFieldValidator,
//     validationAction: "error",
//     validationLevel: "strict",
// });