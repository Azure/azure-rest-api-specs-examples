const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates instance.
 *
 * @summary Creates or updates instance.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Create.json
 */
async function createsOrUpdatesInstance() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const instanceName = "blue";
  const instance = {
    diagnosticStorageProperties: {
      authenticationType: "KeyBased",
      connectionString: "string",
      resourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/adu-resource-group/providers/Microsoft.Storage/storageAccounts/testAccount",
    },
    enableDiagnostics: false,
    iotHubs: [
      {
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Devices/IotHubs/blue-contoso-hub",
      },
    ],
    location: "westus2",
  };
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.instances.beginCreateAndWait(
    resourceGroupName,
    accountName,
    instanceName,
    instance
  );
  console.log(result);
}

createsOrUpdatesInstance().catch(console.error);
