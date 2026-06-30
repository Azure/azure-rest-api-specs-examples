const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 *
 * @summary updates share properties as specified in request body. Properties not mentioned in the request will not be changed. Update fails if the specified share does not already exist.
 * x-ms-original-file: 2026-04-01/FileShareAclsPatch.json
 */
async function updateShareAcls() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.update("res3376", "sto328", "share6185", {
    signedIdentifiers: [
      {
        accessPolicy: {
          expiryTime: new Date("2021-05-01T08:49:37.0000000Z"),
          permission: "rwd",
          startTime: new Date("2021-04-01T08:49:37.0000000Z"),
        },
        id: "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
      },
    ],
  });
  console.log(result);
}
