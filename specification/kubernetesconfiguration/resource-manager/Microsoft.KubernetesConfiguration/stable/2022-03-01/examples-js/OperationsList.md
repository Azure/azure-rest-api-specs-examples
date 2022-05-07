Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-kubernetesconfiguration_5.0.0/sdk/kubernetesconfiguration/arm-kubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the available operations the KubernetesConfiguration resource provider supports.
 *
 * @summary List all the available operations the KubernetesConfiguration resource provider supports.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/OperationsList.json
 */
async function batchAccountDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

batchAccountDelete().catch(console.error);
```
