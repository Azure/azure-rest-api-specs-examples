const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the environment definitions in this project catalog.
 *
 * @summary Lists the environment definitions in this project catalog.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentDefinitions_ListByProjectCatalog.json
 */
async function environmentDefinitionsListByProjectCatalog() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "DevProject";
  const catalogName = "myCatalog";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.environmentDefinitions.listByProjectCatalog(
    resourceGroupName,
    projectName,
    catalogName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
