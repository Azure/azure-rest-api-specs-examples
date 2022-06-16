const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate Heap Dump
 *
 * @summary Generate Heap Dump
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_GenerateHeapDump.json
 */
async function deploymentsGenerateHeapDump() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
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

deploymentsGenerateHeapDump().catch(console.error);
