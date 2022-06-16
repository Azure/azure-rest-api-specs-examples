const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_PowerBI.json
 */
async function createAPowerBiOutput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg7983";
  const jobName = "sj2331";
  const outputName = "output3022";
  const output = {
    datasource: {
      type: "PowerBI",
      dataset: "someDataset",
      groupId: "ac40305e-3e8d-43ac-8161-c33799f43e95",
      groupName: "MyPowerBIGroup",
      refreshToken: "someRefreshToken==",
      table: "someTable",
      tokenUserDisplayName: "Bob Smith",
      tokenUserPrincipalName: "bobsmith@contoso.com",
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

createAPowerBiOutput().catch(console.error);
