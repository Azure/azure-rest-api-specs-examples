const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether instance exists.
 *
 * @summary Checks whether instance exists.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Head.json
 */
async function checksWhetherInstanceExists() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const instanceName = "blue";
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.instances.head(resourceGroupName, accountName, instanceName);
  console.log(result);
}

checksWhetherInstanceExists().catch(console.error);
