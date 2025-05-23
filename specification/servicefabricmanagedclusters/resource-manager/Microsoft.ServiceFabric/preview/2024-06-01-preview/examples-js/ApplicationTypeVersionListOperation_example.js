const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all application type version resources created or in the process of being created in the Service Fabric managed application type name resource.
 *
 * @summary Gets all application type version resources created or in the process of being created in the Service Fabric managed application type name resource.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/ApplicationTypeVersionListOperation_example.json
 */
async function getAListOfApplicationTypeVersionResources() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationTypeVersions.listByApplicationTypes(
    resourceGroupName,
    clusterName,
    applicationTypeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
