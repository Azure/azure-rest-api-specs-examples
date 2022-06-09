```javascript
const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds new or replaces existing certificate.
 *
 * @summary Adds new or replaces existing certificate.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_certificatescreateorupdate.json
 */
async function certificatesCreateOrUpdate() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
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

certificatesCreateOrUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-iothub_6.1.1/sdk/iothub/arm-iothub/README.md) on how to add the SDK to your project and authenticate.
