const { ResourceConnectorManagementClient } = require("@azure/arm-resourceconnector");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an Appliance with the specified Resource Name in the specified Resource Group and Subscription.
 *
 * @summary Updates an Appliance with the specified Resource Name in the specified Resource Group and Subscription.
 * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesPatch.json
 */
async function updateAppliance() {
  const subscriptionId =
    process.env["RESOURCECONNECTOR_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["RESOURCECONNECTOR_RESOURCE_GROUP"] || "testresourcegroup";
  const resourceName = "appliance01";
  const tags = { key: "value" };
  const options = { tags };
  const credential = new DefaultAzureCredential();
  const client = new ResourceConnectorManagementClient(credential, subscriptionId);
  const result = await client.appliances.update(resourceGroupName, resourceName, options);
  console.log(result);
}
