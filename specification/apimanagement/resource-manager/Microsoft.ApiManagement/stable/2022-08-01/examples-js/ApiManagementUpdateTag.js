const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the details of the tag specified by its identifier.
 *
 * @summary Updates the details of the tag specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateTag.json
 */
async function apiManagementUpdateTag() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const tagId = "temptag";
  const ifMatch = "*";
  const parameters = { displayName: "temp tag" };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.update(
    resourceGroupName,
    serviceName,
    tagId,
    ifMatch,
    parameters
  );
  console.log(result);
}
