const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Tag description in scope of API
 *
 * @summary Get Tag description in scope of API
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiTagDescription.json
 */
async function apiManagementGetApiTagDescription() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "59d6bb8f1f7fab13dc67ec9b";
  const tagDescriptionId = "59306a29e4bbd510dc24e5f9";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiTagDescription.get(
    resourceGroupName,
    serviceName,
    apiId,
    tagDescriptionId
  );
  console.log(result);
}

apiManagementGetApiTagDescription().catch(console.error);
