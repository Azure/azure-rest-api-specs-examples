const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds new or replaces existing certificate.
 *
 * @summary Adds new or replaces existing certificate.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_certificatescreateorupdate.json
 */
async function certificatesCreateOrUpdate() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "iothub";
  const certificateName = "cert";
  const certificateDescription = {
    properties: { certificate: "############################################" },
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.certificates.createOrUpdate(
    resourceGroupName,
    resourceName,
    certificateName,
    certificateDescription
  );
  console.log(result);
}
