const { SourceControlConfigurationClient } = require("@azure/arm-kubernetesconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Async Operations, currently in progress, in a cluster
 *
 * @summary List Async Operations, currently in progress, in a cluster
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/ListAsyncOperationStatus.json
 */
async function asyncOperationStatusList() {
  const subscriptionId = "subId1";
  const resourceGroupName = "rg1";
  const clusterRp = "Microsoft.Kubernetes";
  const clusterResourceName = "connectedClusters";
  const clusterName = "clusterName1";
  const credential = new DefaultAzureCredential();
  const client = new SourceControlConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operationStatus.list(
    resourceGroupName,
    clusterRp,
    clusterResourceName,
    clusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

asyncOperationStatusList().catch(console.error);
