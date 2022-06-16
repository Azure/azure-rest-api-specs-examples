const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get Diagnostics Category
 *
 * @summary Description for Get Diagnostics Category
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_GetSiteDiagnosticCategory.json
 */
async function getAppDiagnosticCategory() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const siteName = "SampleApp";
  const diagnosticCategory = "availability";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.getSiteDiagnosticCategory(
    resourceGroupName,
    siteName,
    diagnosticCategory
  );
  console.log(result);
}

getAppDiagnosticCategory().catch(console.error);
