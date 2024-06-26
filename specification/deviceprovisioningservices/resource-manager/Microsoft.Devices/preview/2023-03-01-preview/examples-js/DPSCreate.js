const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the metadata of the provisioning service. The usual pattern to modify a property is to retrieve the provisioning service metadata and security metadata, and then combine them with the modified values in a new body to update the provisioning service.
 *
 * @summary Create or update the metadata of the provisioning service. The usual pattern to modify a property is to retrieve the provisioning service metadata and security metadata, and then combine them with the modified values in a new body to update the provisioning service.
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSCreate.json
 */
async function dpsCreate() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName =
    process.env["DEVICEPROVISIONINGSERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const provisioningServiceName = "myFirstProvisioningService";
  const iotDpsDescription = {
    location: "East US",
    properties: { enableDataResidency: false },
    sku: { name: "S1", capacity: 1 },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.iotDpsResource.beginCreateOrUpdateAndWait(
    resourceGroupName,
    provisioningServiceName,
    iotDpsDescription
  );
  console.log(result);
}
