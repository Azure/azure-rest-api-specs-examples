const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes specific Api Version Set.
 *
 * @summary Deletes specific Api Version Set.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiVersionSet.json
 */
async function apiManagementDeleteApiVersionSet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const versionSetId = "a1";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiVersionSet.delete(
    resourceGroupName,
    serviceName,
    versionSetId,
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteApiVersionSet().catch(console.error);
