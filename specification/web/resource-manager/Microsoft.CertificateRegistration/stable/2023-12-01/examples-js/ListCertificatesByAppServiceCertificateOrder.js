const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for List all certificates associated with a certificate order.
 *
 * @summary Description for List all certificates associated with a certificate order.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/ListCertificatesByAppServiceCertificateOrder.json
 */
async function listCertificatesByAppServiceCertificate() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const certificateOrderName = "SampleCertificateOrderName";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appServiceCertificateOrders.listCertificates(
    resourceGroupName,
    certificateOrderName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
