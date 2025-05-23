const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric managed application resource with the specified name.
 *
 * @summary Create or update a Service Fabric managed application resource with the specified name.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/ApplicationPutOperation_example_max.json
 */
async function putAnApplicationWithMaximumParameters() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const parameters = {
    location: "eastus",
    parameters: { param1: "value1" },
    tags: { a: "b" },
    upgradePolicy: {
      applicationHealthPolicy: {
        considerWarningAsError: true,
        defaultServiceTypeHealthPolicy: {
          maxPercentUnhealthyPartitionsPerService: 0,
          maxPercentUnhealthyReplicasPerPartition: 0,
          maxPercentUnhealthyServices: 0,
        },
        maxPercentUnhealthyDeployedApplications: 0,
        serviceTypeHealthPolicyMap: {
          service1: {
            maxPercentUnhealthyPartitionsPerService: 30,
            maxPercentUnhealthyReplicasPerPartition: 30,
            maxPercentUnhealthyServices: 30,
          },
        },
      },
      forceRestart: false,
      instanceCloseDelayDuration: 600,
      recreateApplication: false,
      rollingUpgradeMonitoringPolicy: {
        failureAction: "Rollback",
        healthCheckRetryTimeout: "00:10:00",
        healthCheckStableDuration: "00:05:00",
        healthCheckWaitDuration: "00:02:00",
        upgradeDomainTimeout: "00:15:00",
        upgradeTimeout: "01:00:00",
      },
      upgradeMode: "UnmonitoredAuto",
      upgradeReplicaSetCheckTimeout: 3600,
    },
    version:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.applications.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    parameters,
  );
  console.log(result);
}
