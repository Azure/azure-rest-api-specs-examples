Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update build service agent pool.
 *
 * @summary Create or update build service agent pool.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildServiceAgentPool_UpdatePut.json
 */
async function buildServiceAgentPoolUpdatePut() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const agentPoolName = "default";
  const agentPoolResource = {
    properties: { poolSize: { name: "S3" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceAgentPool.beginUpdatePutAndWait(
    resourceGroupName,
    serviceName,
    buildServiceName,
    agentPoolName,
    agentPoolResource
  );
  console.log(result);
}

buildServiceAgentPoolUpdatePut().catch(console.error);
```
