const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes a content from Front Door.
 *
 * @summary Removes a content from Front Door.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorPurgeContent.json
 */
async function purgeContentFromFrontDoor() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const contentFilePaths = {
    contentPaths: ["/pictures.aspx", "/pictures/*"],
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginPurgeContentAndWait(
    resourceGroupName,
    frontDoorName,
    contentFilePaths
  );
  console.log(result);
}

purgeContentFromFrontDoor().catch(console.error);
