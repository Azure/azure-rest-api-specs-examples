const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Service Fabric application type version resource created or in the process of being created in the Service Fabric application type name resource.
 *
 * @summary Get a Service Fabric application type version resource created or in the process of being created in the Service Fabric application type name resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeVersionGetOperation_example.json
 */
async function getAnApplicationTypeVersion() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const version = "1.0";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.applicationTypeVersions.get(
    resourceGroupName,
    clusterName,
    applicationTypeName,
    version
  );
  console.log(result);
}

getAnApplicationTypeVersion().catch(console.error);
