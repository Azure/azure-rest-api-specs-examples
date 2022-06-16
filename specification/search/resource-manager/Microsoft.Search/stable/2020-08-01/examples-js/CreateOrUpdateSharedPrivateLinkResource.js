const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initiates the creation or update of a shared private link resource managed by the search service in the given resource group.
 *
 * @summary Initiates the creation or update of a shared private link resource managed by the search service in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/CreateOrUpdateSharedPrivateLinkResource.json
 */
async function sharedPrivateLinkResourceCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const sharedPrivateLinkResourceName = "testResource";
  const sharedPrivateLinkResource = {
    properties: {
      groupId: "blob",
      privateLinkResourceId:
        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName",
      requestMessage: "please approve",
      resourceRegion: undefined,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.sharedPrivateLinkResources.beginCreateOrUpdateAndWait(
    resourceGroupName,
    searchServiceName,
    sharedPrivateLinkResourceName,
    sharedPrivateLinkResource
  );
  console.log(result);
}

sharedPrivateLinkResourceCreateOrUpdate().catch(console.error);
