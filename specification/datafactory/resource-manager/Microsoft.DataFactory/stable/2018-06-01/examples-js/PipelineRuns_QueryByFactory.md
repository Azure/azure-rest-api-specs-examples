Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function pipelineRunsQueryByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const filterParameters = {
    filters: [
      {
        operand: "PipelineName",
        operator: "Equals",
        values: ["examplePipeline"],
      },
    ],
    lastUpdatedAfter: new Date("2018-06-16T00:36:44.3345758Z"),
    lastUpdatedBefore: new Date("2018-06-16T00:49:48.3686473Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.pipelineRuns.queryByFactory(
    resourceGroupName,
    factoryName,
    filterParameters
  );
  console.log(result);
}

pipelineRunsQueryByFactory().catch(console.error);
```
