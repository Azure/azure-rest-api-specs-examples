const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric managed application type version resource created or in the process of being created in the Service Fabric managed application type name resource.
 *
 * @summary Get a Service Fabric managed application type version resource created or in the process of being created in the Service Fabric managed application type name resource.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/ApplicationTypeVersionGetOperation_example.json
 */
async function getAnApplicationTypeVersion() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const version = "1.0";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.applicationTypeVersions.get(
    resourceGroupName,
    clusterName,
    applicationTypeName,
    version,
  );
  console.log(result);
}
