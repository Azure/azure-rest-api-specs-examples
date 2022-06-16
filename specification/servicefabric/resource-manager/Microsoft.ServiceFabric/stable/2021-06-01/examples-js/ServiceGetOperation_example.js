const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric service resource created or in the process of being created in the Service Fabric application resource.
 *
 * @summary Get a Service Fabric service resource created or in the process of being created in the Service Fabric application resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceGetOperation_example.json
 */
async function getAService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const serviceName = "myService";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.services.get(
    resourceGroupName,
    clusterName,
    applicationName,
    serviceName
  );
  console.log(result);
}

getAService().catch(console.error);
