const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a network security perimeter configuration.
 *
 * @summary Gets a network security perimeter configuration.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/NetworkSecurityPerimeterConfigurationsGet.json
 */
async function getAnNspConfigByName() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const nspConfigName = "00000001-2222-3333-4444-111144444444.assoc1";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeterConfigurations.get(
    resourceGroupName,
    searchServiceName,
    nspConfigName,
  );
  console.log(result);
}
