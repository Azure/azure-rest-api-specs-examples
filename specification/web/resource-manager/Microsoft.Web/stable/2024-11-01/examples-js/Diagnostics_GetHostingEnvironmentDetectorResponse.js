const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Get Hosting Environment Detector Response
 *
 * @summary Description for Get Hosting Environment Detector Response
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/Diagnostics_GetHostingEnvironmentDetectorResponse.json
 */
async function getAppServiceEnvironmentDetectorResponses() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName =
    process.env["APPSERVICE_RESOURCE_GROUP"] || "Sample-WestUSResourceGroup";
  const name = "SampleAppServiceEnvironment";
  const detectorName = "runtimeavailability";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.getHostingEnvironmentDetectorResponse(
    resourceGroupName,
    name,
    detectorName,
  );
  console.log(result);
}
