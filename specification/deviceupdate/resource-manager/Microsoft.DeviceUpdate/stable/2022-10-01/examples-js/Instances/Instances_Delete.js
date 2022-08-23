const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes instance.
 *
 * @summary Deletes instance.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Instances/Instances_Delete.json
 */
async function deletesInstance() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const instanceName = "blue";
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.instances.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    instanceName
  );
  console.log(result);
}

deletesInstance().catch(console.error);
