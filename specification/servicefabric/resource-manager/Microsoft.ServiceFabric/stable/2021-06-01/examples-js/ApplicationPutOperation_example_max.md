```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric application resource with the specified name.
 *
 * @summary Create or update a Service Fabric application resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_max.json
 */
async function putAnApplicationWithMaximumParameters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const parameters = {
    maximumNodes: 3,
    metrics: [
      {
        name: "metric1",
        maximumCapacity: 3,
        reservationCapacity: 1,
        totalApplicationCapacity: 5,
      },
    ],
    minimumNodes: 1,
    parameters: { param1: "value1" },
    removeApplicationCapacity: false,
    tags: {},
    typeName: "myAppType",
    typeVersion: "1.0",
    upgradePolicy: {
      applicationHealthPolicy: {
        considerWarningAsError: true,
        defaultServiceTypeHealthPolicy: {
          maxPercentUnhealthyPartitionsPerService: 0,
          maxPercentUnhealthyReplicasPerPartition: 0,
          maxPercentUnhealthyServices: 0,
        },
        maxPercentUnhealthyDeployedApplications: 0,
      },
      forceRestart: false,
      rollingUpgradeMonitoringPolicy: {
        failureAction: "Rollback",
        healthCheckRetryTimeout: "00:10:00",
        healthCheckStableDuration: "00:05:00",
        healthCheckWaitDuration: "00:02:00",
        upgradeDomainTimeout: "1.06:00:00",
        upgradeTimeout: "01:00:00",
      },
      upgradeMode: "Monitored",
      upgradeReplicaSetCheckTimeout: "01:00:00",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applications.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    parameters
  );
  console.log(result);
}

putAnApplicationWithMaximumParameters().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.
