const { AzureSphereManagementClient } = require("@azure/arm-sphere");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Image
 *
 * @summary Create a Image
 * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PutImage.json
 */
async function imageCreateOrUpdate() {
  const subscriptionId =
    process.env["SPHERE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SPHERE_RESOURCE_GROUP"] || "MyResourceGroup1";
  const catalogName = "MyCatalog1";
  const imageName = "00000000-0000-0000-0000-000000000000";
  const resource = { properties: { image: "bXliYXNlNjRzdHJpbmc=" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureSphereManagementClient(credential, subscriptionId);
  const result = await client.images.beginCreateOrUpdateAndWait(
    resourceGroupName,
    catalogName,
    imageName,
    resource,
  );
  console.log(result);
}
