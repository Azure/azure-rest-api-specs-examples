const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Cluster
 *
 * @summary Create a Cluster
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Clusters_CreateOrUpdate.json
 */
async function clustersCreateOrUpdate() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const clusterName = "cluster1";
  const cluster = { clusterSize: 3, sku: { name: "AV20" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.clusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateCloudName,
    clusterName,
    cluster,
  );
  console.log(result);
}
