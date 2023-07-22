const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns list of jobs belonging to a backup vault
 *
 * @summary Returns list of jobs belonging to a backup vault
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/JobCRUD/ListJobs.json
 */
async function getJobs() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "62b829ee-7936-40c9-a1c9-47a93f9f3965";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "BugBash1";
  const vaultName = "BugBashVaultForCCYv11";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.list(resourceGroupName, vaultName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
