const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to adds new or replaces existing certificate.
 *
 * @summary adds new or replaces existing certificate.
 * x-ms-original-file: 2026-03-01-preview/CreateOrReplace_Certificates_With_DeviceRegistryPolicy.json
 */
async function createOrReplaceCertificatesWithDeviceRegistryPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.certificates.createOrUpdate("myResourceGroup", "testHub", "cert", {
    properties: { certificate: "############################################" },
  });
  console.log(result);
}
