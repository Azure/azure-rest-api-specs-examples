Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.1/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function commandFailedResult() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const commandId = "def7b3ea71bd4f7e9d226ddbc0f00ad9";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.getCommandResult(
    resourceGroupName,
    resourceName,
    commandId
  );
  console.log(result);
}

commandFailedResult().catch(console.error);
```
