const { WebSiteManagementClient } = require("@azure/arm-appservice-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all certificates for a subscription.
 *
 * @summary Get all certificates for a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2018-02-01/examples/ListCertificates.json
 */
async function listCertificatesForSubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.certificates.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
