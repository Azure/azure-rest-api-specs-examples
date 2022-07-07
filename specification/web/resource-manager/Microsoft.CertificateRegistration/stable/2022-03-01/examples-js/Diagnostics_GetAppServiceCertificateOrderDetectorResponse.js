const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Microsoft.CertificateRegistration call to get a detector response from App Lens.
 *
 * @summary Description for Microsoft.CertificateRegistration call to get a detector response from App Lens.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
 */
async function getAppServiceCertificateOrderDetectorResponse() {
  const subscriptionId = "5700fc96-77b4-4f8d-afce-c353d8c443bd";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const certificateOrderName = "SampleCertificateOrderName";
  const detectorName = "AutoRenewStatus";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result =
    await client.certificateOrdersDiagnostics.getAppServiceCertificateOrderDetectorResponse(
      resourceGroupName,
      certificateOrderName,
      detectorName
    );
  console.log(result);
}

getAppServiceCertificateOrderDetectorResponse().catch(console.error);
