const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Request to mitigate for a given job
 *
 * @summary Request to mitigate for a given job
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/JobMitigate.json
 */
async function mitigate() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const jobName = "TestJobName1";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const mitigateJobRequest = {
    serialNumberCustomerResolutionMap: {
      testDISK1: "MoveToCleanUpDevice",
      testDISK2: "Resume",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.mitigate(jobName, resourceGroupName, mitigateJobRequest);
  console.log(result);
}
