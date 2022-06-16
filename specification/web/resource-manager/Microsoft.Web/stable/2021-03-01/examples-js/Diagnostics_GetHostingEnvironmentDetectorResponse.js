const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get Hosting Environment Detector Response
 *
 * @summary Description for Get Hosting Environment Detector Response
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_GetHostingEnvironmentDetectorResponse.json
 */
async function getAppServiceEnvironmentDetectorResponses() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const name = "SampleAppServiceEnvironment";
  const detectorName = "runtimeavailability";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.getHostingEnvironmentDetectorResponse(
    resourceGroupName,
    name,
    detectorName
  );
  console.log(result);
}

getAppServiceEnvironmentDetectorResponses().catch(console.error);
