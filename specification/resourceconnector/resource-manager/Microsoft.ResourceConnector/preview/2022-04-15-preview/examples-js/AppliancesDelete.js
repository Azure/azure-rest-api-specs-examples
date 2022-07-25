const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Appliance with the specified Resource Name, Resource Group, and Subscription Id.
 *
 * @summary Deletes an Appliance with the specified Resource Name, Resource Group, and Subscription Id.
 * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesDelete.json
 */
async function deleteAppliance() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "testresourcegroup";
  const resourceName = "appliance01";
  const credential = new DefaultAzureCredential();
  const client = new ResourceConnectorManagementClient(credential, subscriptionId);
  const result = await client.appliances.beginDeleteAndWait(resourceGroupName, resourceName);
  console.log(result);
}

deleteAppliance().catch(console.error);
