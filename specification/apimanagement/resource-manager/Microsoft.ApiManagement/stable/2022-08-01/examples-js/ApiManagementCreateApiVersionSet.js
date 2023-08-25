const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a Api Version Set.
 *
 * @summary Creates or Updates a Api Version Set.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiVersionSet.json
 */
async function apiManagementCreateApiVersionSet() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const versionSetId = "api1";
  const parameters = {
    description: "Version configuration",
    displayName: "api set 1",
    versioningScheme: "Segment",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiVersionSet.createOrUpdate(
    resourceGroupName,
    serviceName,
    versionSetId,
    parameters
  );
  console.log(result);
}
