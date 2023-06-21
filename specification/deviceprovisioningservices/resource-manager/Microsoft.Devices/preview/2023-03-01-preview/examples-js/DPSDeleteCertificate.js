const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified certificate associated with the Provisioning Service
 *
 * @summary Deletes the specified certificate associated with the Provisioning Service
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/preview/2023-03-01-preview/examples/DPSDeleteCertificate.json
 */
async function dpsDeleteCertificate() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName =
    process.env["DEVICEPROVISIONINGSERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const ifMatch = "AAAAAAAADGk=";
  const provisioningServiceName = "myFirstProvisioningService";
  const certificateName = "cert";
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.dpsCertificate.delete(
    resourceGroupName,
    ifMatch,
    provisioningServiceName,
    certificateName
  );
  console.log(result);
}
