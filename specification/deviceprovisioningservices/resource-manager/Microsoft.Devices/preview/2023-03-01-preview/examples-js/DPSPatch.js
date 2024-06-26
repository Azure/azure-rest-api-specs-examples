const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an existing provisioning service's tags. to update other fields use the CreateOrUpdate method
 *
 * @summary Update an existing provisioning service's tags. to update other fields use the CreateOrUpdate method
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSPatch.json
 */
async function dpsPatch() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName =
    process.env["DEVICEPROVISIONINGSERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const provisioningServiceName = "myFirstProvisioningService";
  const provisioningServiceTags = { tags: { foo: "bar" } };
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.iotDpsResource.beginUpdateAndWait(
    resourceGroupName,
    provisioningServiceName,
    provisioningServiceTags
  );
  console.log(result);
}
