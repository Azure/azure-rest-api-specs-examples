const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an application.
 *
 * @summary Update an application.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/Application_Update.json
 */
async function applicationUpdate() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const applicationGroupName = "applicationGroup1";
  const applicationName = "application1";
  const application = {
    description: "des1",
    applicationType: "InBuilt",
    commandLineArguments: "arguments",
    commandLineSetting: "Allow",
    filePath: "path",
    friendlyName: "friendly",
    iconIndex: 1,
    iconPath: "icon",
    msixPackageApplicationId: undefined,
    msixPackageFamilyName: undefined,
    showInPortal: true,
  };
  const options = { application };
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.applications.update(
    resourceGroupName,
    applicationGroupName,
    applicationName,
    options,
  );
  console.log(result);
}
