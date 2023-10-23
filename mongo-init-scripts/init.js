for (let i = 1; i <= 4; i++) {
    db.rooms.insert({
        name: `Room ${i}`,
    });
}
for (let i = 1; i <= 10; i++) {
    db.users.insert({
        username: `jobsity${i}`,
        email: `jobsity${i}@example.com`,
        password: "36707e8543df87c5f639ea729d97eb5e1b9eac205d95"
    });
}
print("db 'jobsity' created and data inserted.");
