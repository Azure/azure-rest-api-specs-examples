const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the certificate.
 *
 * @summary Returns the certificate.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getcertificate.json
 */
async function certificatesGet() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testhub";
  const certificateName = "cert";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.certificates.get(resourceGroupName, resourceName, certificateName);
  console.log(result);
}

certificatesGet().catch(console.error);
