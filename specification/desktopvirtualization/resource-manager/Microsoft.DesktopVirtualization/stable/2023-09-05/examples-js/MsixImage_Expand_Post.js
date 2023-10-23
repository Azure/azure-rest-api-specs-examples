const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Expands and Lists MSIX packages in an Image, given the Image Path.
 *
 * @summary Expands and Lists MSIX packages in an Image, given the Image Path.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/MsixImage_Expand_Post.json
 */
async function msixImageExpand() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const hostPoolName = "hostpool1";
  const msixImageURI = { uri: "imagepath" };
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.msixImages.listExpand(
    resourceGroupName,
    hostPoolName,
    msixImageURI
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
