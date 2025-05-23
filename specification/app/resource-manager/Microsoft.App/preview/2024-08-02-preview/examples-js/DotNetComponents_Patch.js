const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a .NET Component using JSON Merge Patch
 *
 * @summary Patches a .NET Component using JSON Merge Patch
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DotNetComponents_Patch.json
 */
async function patchNetComponent() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "mydotnetcomponent";
  const dotNetComponentEnvelope = {
    componentType: "AspireDashboard",
    configurations: [{ propertyName: "dashboard-theme", value: "dark" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.dotNetComponents.beginUpdateAndWait(
    resourceGroupName,
    environmentName,
    name,
    dotNetComponentEnvelope,
  );
  console.log(result);
}
