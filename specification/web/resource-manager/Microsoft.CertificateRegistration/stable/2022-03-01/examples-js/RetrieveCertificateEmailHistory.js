const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Retrieve email history.
 *
 * @summary Description for Retrieve email history.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/RetrieveCertificateEmailHistory.json
 */
async function retrieveCertificateEmailHistory() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "SampleCertOrder";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceCertificateOrders.retrieveCertificateEmailHistory(
    resourceGroupName,
    name
  );
  console.log(result);
}

retrieveCertificateEmailHistory().catch(console.error);
