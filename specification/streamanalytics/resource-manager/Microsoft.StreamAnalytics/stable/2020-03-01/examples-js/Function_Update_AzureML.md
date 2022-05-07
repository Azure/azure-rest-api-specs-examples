Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing function under an existing streaming job. This can be used to partially update (ie. update one or two properties) a function without affecting the rest the job or function definition.
 *
 * @summary Updates an existing function under an existing streaming job. This can be used to partially update (ie. update one or two properties) a function without affecting the rest the job or function definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_AzureML.json
 */
async function updateAnAzureMlFunction() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg7";
  const jobName = "sj9093";
  const functionName = "function588";
  const functionParam = {
    properties: {
      type: "Scalar",
      binding: { type: "Microsoft.MachineLearning/WebService", batchSize: 5000 },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.functions.update(
    resourceGroupName,
    jobName,
    functionName,
    functionParam
  );
  console.log(result);
}

updateAnAzureMlFunction().catch(console.error);
```
