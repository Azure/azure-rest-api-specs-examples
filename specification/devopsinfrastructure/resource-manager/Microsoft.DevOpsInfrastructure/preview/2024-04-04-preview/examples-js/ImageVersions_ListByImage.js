const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List ImageVersion resources by Image
 *
 * @summary List ImageVersion resources by Image
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ImageVersions_ListByImage.json
 */
async function imageVersionsListByImage() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const resourceGroupName =
    process.env["DEVOPSINFRASTRUCTURE_RESOURCE_GROUP"] || "my-resource-group";
  const imageName = "windows-2022";
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.imageVersions.listByImage(resourceGroupName, imageName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
