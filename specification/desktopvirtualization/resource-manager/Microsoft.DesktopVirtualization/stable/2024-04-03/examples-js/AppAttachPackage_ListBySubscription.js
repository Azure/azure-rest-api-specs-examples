const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List App Attach packages in subscription.
 *
 * @summary List App Attach packages in subscription.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/AppAttachPackage_ListBySubscription.json
 */
async function appAttachPackageListBySubscription() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const filter = "HostPoolName eq 'hostpool1'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appAttachPackageOperations.listBySubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
