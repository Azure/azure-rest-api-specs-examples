const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Association
 *
 * @summary Get a Association
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/AssociationGet.json
 */
async function getAssociation() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "tc1";
  const associationName = "as1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.associationsInterface.get(
    resourceGroupName,
    trafficControllerName,
    associationName
  );
  console.log(result);
}
