// --- Init roles ---

var adminRole = {
    _id: 1,
    name: "admin",
    code: 0,
    isRemovable: false,
    paths: [
        {
            path: "*:*"
        }
    ]
};

var userRole = {
    _id: 2,
    name: "user",
    code: 1,
    isRemovable: false,
    paths: [
        {
            path: "assignedTask:create"
        },
        {
            path: "assignedTask:read"
        },
        {
            path: "assignedTask:update"
        },
        {
            path: "assignedTask:list"
        },
        {
            path: "bookReadTask:*"
        },
        {
            path: "company:read"
        },
        {
            path: "company:list"
        },
        {
            path: "companyInvite:read"
        },
        {
            path: "companyInvite:update"
        },
        {
            path: "companyTask:read"
        },
        {
            path: "companyTask:list"
        },
        {
            path: "movieWatchTask:*"
        },
        {
            path: "multistepTask:*"
        },
        {
            path: "reminder:*"
        },
        {
            path: "repeatableTask:*"
        },
        {
            path: "selfTask:*"
        },
        {
            path: "taskMessage:*"
        },
        {
            path: "taskStep:*"
        }
    ]
};

var companyOwnerRole = {
    _id: 3,
    name: "companyOwner",
    code: 2,
    isRemovable: true,
    paths: [
        {
            path: "assignedTask:*"
        },
        {
            path: "company:update"
        },
        {
            path: "companyInvite:*"
        },
        {
            path: "companyTask:*"
        }
    ]
};

var companyManagerRole = {
    _id: 4,
    name: "companyManager",
    code: 3,
    isRemovable: true,
    paths: [
        {
            path: "assignedTask:*"
        },
        {
            path: "company:*"
        },
        {
            path: "companyInvite:*"
        },
        {
            path: "companyTask:*"
        }
    ]
};

var adminRoleId = db.roles.insertOne(adminRole).insertedId;
var userRoleId = db.roles.insertOne(userRole).insertedId;
var companyOwnerId = db.roles.insertOne(companyOwnerRole).insertedId;
var companyManagerId = db.roles.insertOne(companyManagerRole).insertedId;


// --- Init users ---

var admin = {
    nickname: "root",
    email: "admin@email.com",
    password: "$2y$10$SJDNZpt3PJZbY2XqUdJ4Zuk8BlDybt8uLgKwu6dZ2LwhecL1wr3hu",
    gender: "M",
    roles: [{role_id: adminRoleId}],
    tokens: []
};

db.users.insert(admin);