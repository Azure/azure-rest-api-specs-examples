const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric application resource with the specified name.
 *
 * @summary Create or update a Service Fabric application resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_min.json
 */
async function putAnApplicationWithMinimumParameters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const parameters = {
    location: "eastus",
    removeApplicationCapacity: false,
    tags: {},
    typeName: "myAppType",
    typeVersion: "1.0",
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

putAnApplicationWithMinimumParameters().catch(console.error);
