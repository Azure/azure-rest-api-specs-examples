const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified private link resource for the given provisioning service
 *
 * @summary Get the specified private link resource for the given provisioning service
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSGetPrivateLinkResources.json
 */
async function privateLinkResourcesList() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName =
    process.env["DEVICEPROVISIONINGSERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myFirstProvisioningService";
  const groupId = "iotDps";
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.iotDpsResource.getPrivateLinkResources(
    resourceGroupName,
    resourceName,
    groupId
  );
  console.log(result);
}
