const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Association resources by TrafficController
 *
 * @summary List Association resources by TrafficController
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2023-11-01/examples/AssociationsGet.json
 */
async function getAssociations() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "tc1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.associationsInterface.listByTrafficController(
    resourceGroupName,
    trafficControllerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
