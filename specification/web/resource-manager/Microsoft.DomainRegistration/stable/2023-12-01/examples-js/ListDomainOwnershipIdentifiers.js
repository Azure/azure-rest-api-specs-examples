const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Lists domain ownership identifiers.
 *
 * @summary Description for Lists domain ownership identifiers.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-12-01/examples/ListDomainOwnershipIdentifiers.json
 */
async function listDomainOwnershipIdentifiers() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const domainName = "example.com";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.domains.listOwnershipIdentifiers(resourceGroupName, domainName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
