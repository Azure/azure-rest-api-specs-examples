const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_DocumentDB.json
 */
async function createADocumentDbOutput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg7983";
  const jobName = "sj2331";
  const outputName = "output3022";
  const output = {
    datasource: {
      type: "Microsoft.Storage/DocumentDB",
      accountId: "someAccountId",
      accountKey: "accountKey==",
      collectionNamePattern: "collection",
      database: "db01",
      documentId: "documentId",
      partitionKey: "key",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.createOrReplace(
    resourceGroupName,
    jobName,
    outputName,
    output
  );
  console.log(result);
}

createADocumentDbOutput().catch(console.error);
