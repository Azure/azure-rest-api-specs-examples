const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initiates the deletion of the shared private link resource from the search service.
 *
 * @summary Initiates the deletion of the shared private link resource from the search service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/DeleteSharedPrivateLinkResource.json
 */
async function sharedPrivateLinkResourceDelete() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const sharedPrivateLinkResourceName = "testResource";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.sharedPrivateLinkResources.beginDeleteAndWait(
    resourceGroupName,
    searchServiceName,
    sharedPrivateLinkResourceName,
  );
  console.log(result);
}
