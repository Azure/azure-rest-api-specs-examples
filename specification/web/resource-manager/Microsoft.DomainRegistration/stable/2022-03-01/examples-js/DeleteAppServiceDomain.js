const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Delete a domain.
 *
 * @summary Description for Delete a domain.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/DeleteAppServiceDomain.json
 */
async function deleteAppServiceDomain() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const domainName = "example.com";
  const forceHardDeleteDomain = true;
  const options = { forceHardDeleteDomain };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.domains.delete(resourceGroupName, domainName, options);
  console.log(result);
}

deleteAppServiceDomain().catch(console.error);
