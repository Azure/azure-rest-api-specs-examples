const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Traffic Controller Association
 *
 * @summary Update a Traffic Controller Association
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationPatch.json
 */
async function updateAssociation() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "TC1";
  const associationName = "associatedvnet-1";
  const properties = {
    properties: { associationType: "subnets", subnet: { id: "subnetFullRef" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.associationsInterface.update(
    resourceGroupName,
    trafficControllerName,
    associationName,
    properties
  );
  console.log(result);
}
