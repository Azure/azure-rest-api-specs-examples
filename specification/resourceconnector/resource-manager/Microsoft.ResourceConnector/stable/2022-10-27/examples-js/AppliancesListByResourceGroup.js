const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Appliances in the specified subscription and resource group. The operation returns properties of each Appliance.
 *
 * @summary Gets a list of Appliances in the specified subscription and resource group. The operation returns properties of each Appliance.
 * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListByResourceGroup.json
 */
async function listAppliancesByResourceGroup() {
  const subscriptionId =
    process.env["RESOURCECONNECTOR_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["RESOURCECONNECTOR_RESOURCE_GROUP"] || "testresourcegroup";
  const credential = new DefaultAzureCredential();
  const client = new ResourceConnectorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appliances.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
