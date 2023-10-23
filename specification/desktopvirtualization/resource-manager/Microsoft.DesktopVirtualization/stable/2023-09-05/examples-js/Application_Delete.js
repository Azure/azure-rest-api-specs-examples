const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove an application.
 *
 * @summary Remove an application.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/Application_Delete.json
 */
async function applicationDelete() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const applicationGroupName = "applicationGroup1";
  const applicationName = "application1";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.applications.delete(
    resourceGroupName,
    applicationGroupName,
    applicationName
  );
  console.log(result);
}
