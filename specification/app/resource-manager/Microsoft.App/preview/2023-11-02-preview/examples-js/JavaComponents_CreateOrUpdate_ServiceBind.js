const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Java Component in a Managed Environment.
 *
 * @summary Creates or updates a Java Component in a Managed Environment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/JavaComponents_CreateOrUpdate_ServiceBind.json
 */
async function createOrUpdateJavaComponentWithServiceBinds() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "myjavacomponent";
  const javaComponentEnvelope = {
    componentType: "SpringBootAdmin",
    configurations: [
      { propertyName: "spring.boot.admin.ui.enable-toasts", value: "true" },
      {
        propertyName: "spring.boot.admin.monitor.status-interval",
        value: "10000ms",
      },
    ],
    serviceBinds: [
      {
        name: "yellowcat",
        serviceId:
          "/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/yellowcat",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.javaComponents.beginCreateOrUpdateAndWait(
    resourceGroupName,
    environmentName,
    name,
    javaComponentEnvelope,
  );
  console.log(result);
}
