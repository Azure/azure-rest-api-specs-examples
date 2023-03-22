const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Transfer out domain to another registrar
 *
 * @summary Transfer out domain to another registrar
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-09-01/examples/TransferOutDomain.json
 */
async function transferOutDomain() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const domainName = "example.com";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.domains.transferOut(resourceGroupName, domainName);
  console.log(result);
}
