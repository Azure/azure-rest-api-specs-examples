const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Delete an existing certificate order.
 *
 * @summary Description for Delete an existing certificate order.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/DeleteAppServiceCertificateOrder.json
 */
async function deleteAppServiceCertificateOrder() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const certificateOrderName = "SampleCertificateOrderName";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceCertificateOrders.delete(
    resourceGroupName,
    certificateOrderName
  );
  console.log(result);
}

deleteAppServiceCertificateOrder().catch(console.error);
