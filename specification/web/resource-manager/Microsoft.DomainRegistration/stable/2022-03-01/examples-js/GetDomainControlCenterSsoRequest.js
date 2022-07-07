const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Generate a single sign-on request for the domain management portal.
 *
 * @summary Description for Generate a single sign-on request for the domain management portal.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/GetDomainControlCenterSsoRequest.json
 */
async function getDomainControlCenterSsoRequest() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.domains.getControlCenterSsoRequest();
  console.log(result);
}

getDomainControlCenterSsoRequest().catch(console.error);
