const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if resource name is available.
 *
 * @summary Checks if resource name is available.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ContainerApps_CheckNameAvailability.json
 */
async function containerAppsCheckNameAvailability() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "testcontainerenv";
  const checkNameAvailabilityRequest = {
    name: "testcappname",
    type: "Microsoft.App/containerApps",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.namespaces.checkNameAvailability(
    resourceGroupName,
    environmentName,
    checkNameAvailabilityRequest,
  );
  console.log(result);
}
