const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Service Fabric service resource with the specified name.
 *
 * @summary Update a Service Fabric service resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePatchOperation_example.json
 */
async function patchAService() {
  const subscriptionId =
    process.env["SERVICEFABRIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRIC_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const serviceName = "myService";
  const parameters = {
    serviceKind: "Stateless",
    serviceLoadMetrics: [{ name: "metric1", weight: "Low" }],
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.services.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    serviceName,
    parameters,
  );
  console.log(result);
}
