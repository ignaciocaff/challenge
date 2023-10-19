for (let i = 1; i <= 8; i++) {
    db.rooms.insert({
        name: `Room ${i}`,
    });
}
for (let i = 1; i <= 10; i++) {
    db.users.insert({
        username: `jobsity${i}`,
        email: `jobsity${i}@example.com`,
        password: "123456"
    });
}
print("db 'jobsity' created and data inserted.");
