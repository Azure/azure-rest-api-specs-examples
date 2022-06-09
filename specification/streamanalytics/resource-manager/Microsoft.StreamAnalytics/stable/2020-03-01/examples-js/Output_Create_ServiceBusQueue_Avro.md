```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_ServiceBusQueue_Avro.json
 */
async function createAServiceBusQueueOutputWithAvroSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3410";
  const jobName = "sj5095";
  const outputName = "output3456";
  const output = {
    datasource: {
      type: "Microsoft.ServiceBus/Queue",
      propertyColumns: ["column1", "column2"],
      queueName: "sdkqueue",
      serviceBusNamespace: "sdktest",
      sharedAccessPolicyKey: "sharedAccessPolicyKey=",
      sharedAccessPolicyName: "RootManageSharedAccessKey",
      systemPropertyColumns: { MessageId: "col3", PartitionKey: "col4" },
    },
    serialization: { type: "Avro" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.createOrReplace(
    resourceGroupName,
    jobName,
    outputName,
    output
  );
  console.log(result);
}

createAServiceBusQueueOutputWithAvroSerialization().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
