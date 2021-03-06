const firebase = require("@firebase/testing");
import fs from "fs";

const projectId = "firestore-emulator-example";
const firebasePort = require("../../firebase.json").emulators.firestore.port;
const port = firebasePort /** Exists? */ ? firebasePort : 8080;
const coverageUrl = `http://localhost:${port}/emulator/v1/projects/${projectId}:ruleCoverage.html`;

const rules = fs.readFileSync("firestore.rules", "utf8");

function authedApp(auth) {
  return firebase.initializeTestApp({ projectId, auth }).firestore();
}

beforeEach(async () => {
  // Clear the database between tests
  await firebase.clearFirestoreData({ projectId });
});

beforeAll(async () => {
  await firebase.loadFirestoreRules({ projectId, rules });
});

afterAll(async () => {
  await Promise.all(firebase.apps().map(app => app.delete()));
  console.log(`View rule coverage information at ${coverageUrl}\n`);
});

describe("firestore.rules", () => {
  it("require users to log in before creating a profile", async () => {
    const db = authedApp(null);
    const profile = db.collection("users").doc("alice");
    await firebase.assertFails(profile.set({ birthday: "January 1" }));
  });

  it("should only let users create their own profile", async () => {
    const db = authedApp({ uid: "alice" });
    await firebase.assertSucceeds(
      db
        .collection("users")
        .doc("alice")
        .set({
          birthday: "January 1",
          createdAt: firebase.firestore.FieldValue.serverTimestamp()
        })
    );
    await firebase.assertFails(
      db
        .collection("users")
        .doc("bob")
        .set({
          birthday: "January 1",
          createdAt: firebase.firestore.FieldValue.serverTimestamp()
        })
    );
  });

  it("should allow admins to write for others", async () => {
    const db = authedApp({ uid: "alice", role: "admin" });
    await firebase.assertSucceeds(
      db
        .collection("users")
        .doc("bob")
        .set({
          birthday: "January 1",
          createdAt: firebase.firestore.FieldValue.serverTimestamp()
        })
    );
  });

  it("should allow admins to write Websites", async () => {
    const db = authedApp({ uid: "alice", role: "admin" });
    await firebase.assertSucceeds(
      db
        .collection("Websites")
        .doc("bbb")
        .set({
          bbb: "ccc"
        })
    );
  });

  it("requires users to log in before accessing documents", async () => {
    const db = authedApp(null);
    const entity = db.collection("entities").doc("alice");
    await firebase.assertFails(entity.get());
  });

  it("users cannot modify most documents in the database", async () => {
    const db = authedApp(null);
    const entity = db.collection("entities").doc("alice");
    await firebase.assertFails(entity.set({ b: "c" }));
  });
});
