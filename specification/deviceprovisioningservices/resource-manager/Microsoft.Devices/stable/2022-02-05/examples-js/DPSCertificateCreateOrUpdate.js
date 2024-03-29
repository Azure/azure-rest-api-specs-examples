const { IotDpsClient } = require("@azure/arm-deviceprovisioningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add new certificate or update an existing certificate.
 *
 * @summary Add new certificate or update an existing certificate.
 * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCertificateCreateOrUpdate.json
 */
async function dpsCreateOrUpdateCertificate() {
  const subscriptionId =
    process.env["DEVICEPROVISIONINGSERVICES_SUBSCRIPTION_ID"] ||
    "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName =
    process.env["DEVICEPROVISIONINGSERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const provisioningServiceName = "myFirstProvisioningService";
  const certificateName = "cert";
  const certificateDescription = {
    properties: {
      certificate: Buffer.from("############################################"),
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new IotDpsClient(credential, subscriptionId);
  const result = await client.dpsCertificate.createOrUpdate(
    resourceGroupName,
    provisioningServiceName,
    certificateName,
    certificateDescription
  );
  console.log(result);
}
