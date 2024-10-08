const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an application.
 *
 * @summary Create or update an application.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/Application_Create.json
 */
async function applicationCreate() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const applicationGroupName = "applicationGroup1";
  const applicationName = "application1";
  const application = {
    description: "des1",
    commandLineArguments: "arguments",
    commandLineSetting: "Allow",
    filePath: "path",
    friendlyName: "friendly",
    iconIndex: 1,
    iconPath: "icon",
    showInPortal: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.applications.createOrUpdate(
    resourceGroupName,
    applicationGroupName,
    applicationName,
    application,
  );
  console.log(result);
}
