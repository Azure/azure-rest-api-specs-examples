const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Reissue an existing certificate order.
 *
 * @summary Description for Reissue an existing certificate order.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/ReissueAppServiceCertificateOrder.json
 */
async function reissueAppServiceCertificateOrder() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const certificateOrderName = "SampleCertificateOrderName";
  const reissueCertificateOrderRequest = {
    csr: "CSR1223238Value",
    delayExistingRevokeInHours: 2,
    isPrivateKeyExternal: false,
    keySize: 2048,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceCertificateOrders.reissue(
    resourceGroupName,
    certificateOrderName,
    reissueCertificateOrderRequest
  );
  console.log(result);
}

reissueAppServiceCertificateOrder().catch(console.error);
