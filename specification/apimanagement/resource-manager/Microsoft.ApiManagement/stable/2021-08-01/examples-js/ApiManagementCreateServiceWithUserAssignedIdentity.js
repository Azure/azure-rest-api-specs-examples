const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateServiceWithUserAssignedIdentity.json
 */
async function apiManagementCreateServiceWithUserAssignedIdentity() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const parameters = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subid/resourcegroups/rg1/providers/MicrosoftManagedIdentity/userAssignedIdentities/apimService1":
          {},
      },
    },
    location: "West US",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    sku: { name: "Consumption", capacity: 0 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}

apiManagementCreateServiceWithUserAssignedIdentity().catch(console.error);
