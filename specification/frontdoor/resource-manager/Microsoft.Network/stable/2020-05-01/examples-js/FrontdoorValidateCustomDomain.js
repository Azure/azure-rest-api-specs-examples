const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates the custom domain mapping to ensure it maps to the correct Front Door endpoint in DNS.
 *
 * @summary Validates the custom domain mapping to ensure it maps to the correct Front Door endpoint in DNS.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorValidateCustomDomain.json
 */
async function frontDoorValidateCustomDomain() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const customDomainProperties = {
    hostName: "www.someDomain.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontDoors.validateCustomDomain(
    resourceGroupName,
    frontDoorName,
    customDomainProperties
  );
  console.log(result);
}

frontDoorValidateCustomDomain().catch(console.error);
