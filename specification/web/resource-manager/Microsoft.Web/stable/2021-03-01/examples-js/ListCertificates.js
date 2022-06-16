const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get all certificates for a subscription.
 *
 * @summary Description for Get all certificates for a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListCertificates.json
 */
async function listCertificatesForSubscription() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.certificates.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCertificatesForSubscription().catch(console.error);
