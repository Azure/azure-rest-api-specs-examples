const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a diagnostics result of a Container App.
 *
 * @summary Get a diagnostics result of a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ContainerAppsDiagnostics_Get.json
 */
async function getContainerAppDiagnosticsInfo() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "f07f3711-b45e-40fe-a941-4e6d93f851e6";
  const resourceGroupName =
    process.env["APPCONTAINERS_RESOURCE_GROUP"] || "mikono-workerapp-test-rg";
  const containerAppName = "mikono-capp-stage1";
  const detectorName = "cappcontainerappnetworkIO";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsDiagnostics.getDetector(
    resourceGroupName,
    containerAppName,
    detectorName,
  );
  console.log(result);
}
