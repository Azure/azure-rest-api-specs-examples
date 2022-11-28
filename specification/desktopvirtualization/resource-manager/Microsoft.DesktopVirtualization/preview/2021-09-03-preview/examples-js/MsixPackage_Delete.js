const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove an MSIX Package.
 *
 * @summary Remove an MSIX Package.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixPackage_Delete.json
 */
async function msixPackageDelete() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = "resourceGroup1";
  const hostPoolName = "hostpool1";
  const msixPackageFullName = "packagefullname";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.msixPackages.delete(
    resourceGroupName,
    hostPoolName,
    msixPackageFullName
  );
  console.log(result);
}

msixPackageDelete().catch(console.error);
