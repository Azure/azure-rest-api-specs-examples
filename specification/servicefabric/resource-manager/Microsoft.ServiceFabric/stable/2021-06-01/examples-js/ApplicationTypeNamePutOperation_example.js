const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric application type name resource with the specified name.
 *
 * @summary Create or update a Service Fabric application type name resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeNamePutOperation_example.json
 */
async function putAnApplicationType() {
  const subscriptionId =
    process.env["SERVICEFABRIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRIC_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const parameters = { tags: {} };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applicationTypes.createOrUpdate(
    resourceGroupName,
    clusterName,
    applicationTypeName,
    parameters,
  );
  console.log(result);
}
