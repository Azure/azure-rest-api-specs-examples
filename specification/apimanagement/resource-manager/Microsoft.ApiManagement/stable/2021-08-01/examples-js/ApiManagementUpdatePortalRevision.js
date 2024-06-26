const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the description of specified portal revision or makes it current.
 *
 * @summary Updates the description of specified portal revision or makes it current.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdatePortalRevision.json
 */
async function apiManagementUpdatePortalRevision() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const portalRevisionId = "20201112101010";
  const ifMatch = "*";
  const parameters = {
    description: "portal revision update",
    isCurrent: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.portalRevision.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    portalRevisionId,
    ifMatch,
    parameters
  );
  console.log(result);
}
