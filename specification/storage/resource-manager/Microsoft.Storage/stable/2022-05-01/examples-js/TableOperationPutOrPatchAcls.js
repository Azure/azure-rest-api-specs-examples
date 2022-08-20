const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new table with the specified table name, under the specified account.
 *
 * @summary Creates a new table with the specified table name, under the specified account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/TableOperationPutOrPatchAcls.json
 */
async function tableOperationPutOrPatchAcls() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const tableName = "table6185";
  const parameters = {
    signedIdentifiers: [
      {
        accessPolicy: {
          expiryTime: new Date("2022-03-20T08:49:37.0000000Z"),
          permission: "raud",
          startTime: new Date("2022-03-17T08:49:37.0000000Z"),
        },
        id: "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
      },
      {
        accessPolicy: {
          expiryTime: new Date("2022-03-20T08:49:37.0000000Z"),
          permission: "rad",
          startTime: new Date("2022-03-17T08:49:37.0000000Z"),
        },
        id: "PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI",
      },
    ],
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.tableOperations.create(
    resourceGroupName,
    accountName,
    tableName,
    options
  );
  console.log(result);
}

tableOperationPutOrPatchAcls().catch(console.error);
