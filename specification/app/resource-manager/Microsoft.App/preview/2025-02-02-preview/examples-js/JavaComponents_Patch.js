const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Patches a Java Component using JSON Merge Patch
 *
 * @summary Patches a Java Component using JSON Merge Patch
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/JavaComponents_Patch.json
 */
async function patchJavaComponent() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "myjavacomponent";
  const javaComponentEnvelope = {
    properties: {
      componentType: "SpringBootAdmin",
      configurations: [
        { propertyName: "spring.boot.admin.ui.enable-toasts", value: "true" },
        {
          propertyName: "spring.boot.admin.monitor.status-interval",
          value: "10000ms",
        },
      ],
      scale: { maxReplicas: 1, minReplicas: 1 },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.javaComponents.beginUpdateAndWait(
    resourceGroupName,
    environmentName,
    name,
    javaComponentEnvelope,
  );
  console.log(result);
}
