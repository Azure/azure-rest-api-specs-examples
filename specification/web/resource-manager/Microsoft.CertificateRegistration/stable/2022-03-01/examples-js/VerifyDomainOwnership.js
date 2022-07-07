const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Verify domain ownership for this certificate order.
 *
 * @summary Description for Verify domain ownership for this certificate order.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/VerifyDomainOwnership.json
 */
async function verifyDomainOwnership() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const certificateOrderName = "SampleCertificateOrderName";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceCertificateOrders.verifyDomainOwnership(
    resourceGroupName,
    certificateOrderName
  );
  console.log(result);
}

verifyDomainOwnership().catch(console.error);
