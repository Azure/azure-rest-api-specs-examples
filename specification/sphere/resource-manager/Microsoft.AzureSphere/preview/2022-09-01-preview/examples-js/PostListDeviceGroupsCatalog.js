const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the device groups for the catalog.
 *
 * @summary List the device groups for the catalog.
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeviceGroupsCatalog.json
 */
async function catalogsListDeviceGroups() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const listDeviceGroupsRequest = {
    deviceGroupName: "MyDeviceGroup1",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.catalogs.listDeviceGroups(
    resourceGroupName,
    catalogName,
    listDeviceGroupsRequest
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
