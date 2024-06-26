const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all application type version resources created or in the process of being created in the Service Fabric application type name resource.
 *
 * @summary Gets all application type version resources created or in the process of being created in the Service Fabric application type name resource.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeVersionListOperation_example.json
 */
async function getAListOfApplicationTypeVersionResources() {
  const subscriptionId =
    process.env["SERVICEFABRIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRIC_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationTypeName = "myAppType";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationTypeVersions.list(
    resourceGroupName,
    clusterName,
    applicationTypeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
