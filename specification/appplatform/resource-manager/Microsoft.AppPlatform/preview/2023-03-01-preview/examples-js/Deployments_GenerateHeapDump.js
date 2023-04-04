const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate Heap Dump
 *
 * @summary Generate Heap Dump
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Deployments_GenerateHeapDump.json
 */
async function deploymentsGenerateHeapDump() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const deploymentName = "mydeployment";
  const diagnosticParameters = {
    appInstance: "myappinstance",
    filePath: "/byos/diagnose",
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginGenerateHeapDumpAndWait(
    resourceGroupName,
    serviceName,
    appName,
    deploymentName,
    diagnosticParameters
  );
  console.log(result);
}
