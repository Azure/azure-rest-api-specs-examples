const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Wiki for an API or updates an existing one.
 *
 * @summary Creates a new Wiki for an API or updates an existing one.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiWiki.json
 */
async function apiManagementCreateApiWiki() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "57d1f7558aa04f15146d9d8a";
  const parameters = {
    documents: [{ documentationId: "docId1" }, { documentationId: "docId2" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiWiki.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    parameters
  );
  console.log(result);
}
