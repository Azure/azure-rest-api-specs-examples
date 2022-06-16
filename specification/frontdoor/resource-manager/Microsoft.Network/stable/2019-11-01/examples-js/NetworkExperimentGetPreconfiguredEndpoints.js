const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Preconfigured Endpoints
 *
 * @summary Gets a list of Preconfigured Endpoints
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetPreconfiguredEndpoints.json
 */
async function getsAListOfPreconfiguredEndpoints() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.preconfiguredEndpoints.list(resourceGroupName, profileName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsAListOfPreconfiguredEndpoints().catch(console.error);
