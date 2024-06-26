const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a resource group.
 *
 * @summary Creates or updates a resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/CreateResourceGroup.json
 */
async function createOrUpdateAResourceGroup() {
  const subscriptionId = process.env["RESOURCES_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "myResourceGroup";
  const parameters = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.resourceGroups.createOrUpdate(resourceGroupName, parameters);
  console.log(result);
}
