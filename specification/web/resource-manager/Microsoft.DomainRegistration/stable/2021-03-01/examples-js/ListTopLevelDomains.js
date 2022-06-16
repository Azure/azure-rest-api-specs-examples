const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get all top-level domains supported for registration.
 *
 * @summary Description for Get all top-level domains supported for registration.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-03-01/examples/ListTopLevelDomains.json
 */
async function listTopLevelDomains() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topLevelDomains.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTopLevelDomains().catch(console.error);
