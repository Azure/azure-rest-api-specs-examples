const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for List all certificate orders in a subscription.
 *
 * @summary Description for List all certificate orders in a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-09-01/examples/ListAppServiceCertificateOrdersBySubscription.json
 */
async function listAppServiceCertificateOrdersBySubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appServiceCertificateOrders.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
