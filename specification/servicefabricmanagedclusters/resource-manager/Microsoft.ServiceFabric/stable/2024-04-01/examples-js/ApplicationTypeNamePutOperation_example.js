const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric managed application type name resource with the specified name.
 *
 * @summary Create or update a Service Fabric managed application type name resource with the specified name.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ApplicationTypeNamePutOperation_example.json
 */
async function putAnApplicationType() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const parameters = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.applicationTypes.createOrUpdate(
    resourceGroupName,
    clusterName,
    applicationTypeName,
    parameters,
  );
  console.log(result);
}
