const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a group.
 *
 * @summary Creates or Updates a group.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGroupExternal.json
 */
async function apiManagementCreateGroupExternal() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const groupId = "aadGroup";
  const parameters = {
    type: "external",
    description: "new group to test",
    displayName: "NewGroup (samiraad.onmicrosoft.com)",
    externalId: "aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.group.createOrUpdate(
    resourceGroupName,
    serviceName,
    groupId,
    parameters
  );
  console.log(result);
}

apiManagementCreateGroupExternal().catch(console.error);
