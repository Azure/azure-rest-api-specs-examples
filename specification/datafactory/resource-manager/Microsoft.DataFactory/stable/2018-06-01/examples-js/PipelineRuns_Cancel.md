Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.6.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancel a pipeline run by its run ID.
 *
 * @summary Cancel a pipeline run by its run ID.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.json
 */
async function pipelineRunsCancel() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const runId = "16ac5348-ff82-4f95-a80d-638c1d47b721";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.pipelineRuns.cancel(resourceGroupName, factoryName, runId);
  console.log(result);
}

pipelineRunsCancel().catch(console.error);
```
