const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Traffic Controller Association
 *
 * @summary Create a Traffic Controller Association
 * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationPut.json
 */
async function putAssociation() {
  const subscriptionId = process.env["SERVICENETWORKING_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SERVICENETWORKING_RESOURCE_GROUP"] || "rg1";
  const trafficControllerName = "TC1";
  const associationName = "associatedvnet-1";
  const resource = {
    associationType: "subnets",
    location: "West US",
    subnet: { id: "subnetFullRef" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.associationsInterface.beginCreateOrUpdateAndWait(
    resourceGroupName,
    trafficControllerName,
    associationName,
    resource
  );
  console.log(result);
}
