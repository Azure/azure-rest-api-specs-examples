const { MachineLearningComputeManagementClient } = require("@azure/arm-machinelearningcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an operationalization cluster.
 *
 * @summary Create or update an operationalization cluster.
 * x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/OperationalizationClusters_CreateOrUpdate.json
 */
async function putOperationalizationCluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const clusterName = "myCluster";
  const parameters = {
    description: "My Operationalization Cluster",
    clusterType: "ACS",
    containerService: {
      orchestratorProperties: {
        servicePrincipal: {
          clientId: "abcdefghijklmnopqrt",
          secret: "<secret>",
        },
      },
      orchestratorType: "Kubernetes",
    },
    globalServiceConfiguration: {
      ssl: {
        cert: "afjdklq2131casfakld=",
        cname: "foo.bar.com",
        key: "flksdafkldsajf=",
        status: "Enabled",
      },
    },
    location: "West US",
    tags: { key1: "alpha", key2: "beta" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningComputeManagementClient(credential, subscriptionId);
  const result = await client.operationalizationClusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

putOperationalizationCluster().catch(console.error);
