```javascript
const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new metric configuration or updates an existing one for a role.
 *
 * @summary Creates a new metric configuration or updates an existing one for a role.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/PutMonitoringConfig.json
 */
async function putMonitoringConfig() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const roleName = "testrole";
  const resourceGroupName = "GroupForEdgeAutomation";
  const monitoringMetricConfiguration = {
    metricConfigurations: [
      {
        counterSets: [{ counters: [{ name: "test" }] }],
        mdmAccount: "test",
        metricNameSpace: "test",
        resourceId: "test",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.monitoringConfig.beginCreateOrUpdateAndWait(
    deviceName,
    roleName,
    resourceGroupName,
    monitoringMetricConfiguration
  );
  console.log(result);
}

putMonitoringConfig().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-databoxedge_2.0.1/sdk/databoxedge/arm-databoxedge/README.md) on how to add the SDK to your project and authenticate.
