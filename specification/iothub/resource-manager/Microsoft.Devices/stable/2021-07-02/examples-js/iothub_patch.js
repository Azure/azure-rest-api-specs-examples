const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an existing IoT Hub tags. to update other fields use the CreateOrUpdate method
 *
 * @summary Update an existing IoT Hub tags. to update other fields use the CreateOrUpdate method
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_patch.json
 */
async function iotHubResourceUpdate() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myHub";
  const iotHubTags = { tags: { foo: "bar" } };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    iotHubTags
  );
  console.log(result);
}

iotHubResourceUpdate().catch(console.error);
