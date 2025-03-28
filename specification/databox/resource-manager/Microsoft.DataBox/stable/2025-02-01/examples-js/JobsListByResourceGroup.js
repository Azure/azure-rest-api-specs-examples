const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all the jobs available under the given resource group.
 *
 * @summary Lists all the jobs available under the given resource group.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/JobsListByResourceGroup.json
 */
async function jobsListByResourceGroup() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.jobs.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
