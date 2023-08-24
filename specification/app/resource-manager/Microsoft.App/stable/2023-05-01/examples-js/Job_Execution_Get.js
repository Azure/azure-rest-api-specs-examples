const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of a single job execution
 *
 * @summary Get details of a single job execution
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Job_Execution_Get.json
 */
async function getASingleJobExecution() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const jobName = "testcontainerAppsJob0";
  const jobExecutionName = "jobExecution1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobExecution(resourceGroupName, jobName, jobExecutionName);
  console.log(result);
}
