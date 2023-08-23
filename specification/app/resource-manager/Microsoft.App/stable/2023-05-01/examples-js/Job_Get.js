const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the properties of a Container Apps Job.
 *
 * @summary Get the properties of a Container Apps Job.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Job_Get.json
 */
async function getContainerAppsJob() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const jobName = "testcontainerAppsJob0";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobs.get(resourceGroupName, jobName);
  console.log(result);
}
