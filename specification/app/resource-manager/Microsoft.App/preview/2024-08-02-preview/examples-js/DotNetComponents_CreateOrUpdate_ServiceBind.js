const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a .NET Component in a Managed Environment.
 *
 * @summary Creates or updates a .NET Component in a Managed Environment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DotNetComponents_CreateOrUpdate_ServiceBind.json
 */
async function createOrUpdateNetComponentWithServiceBinds() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "mydotnetcomponent";
  const dotNetComponentEnvelope = {
    componentType: "AspireDashboard",
    configurations: [{ propertyName: "dashboard-theme", value: "dark" }],
    serviceBinds: [
      {
        name: "yellowcat",
        serviceId:
          "/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/dotNetComponents/yellowcat",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.dotNetComponents.beginCreateOrUpdateAndWait(
    resourceGroupName,
    environmentName,
    name,
    dotNetComponentEnvelope,
  );
  console.log(result);
}
