const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Verifies if this VNET is compatible with an App Service Environment by analyzing the Network Security Group rules.
 *
 * @summary Description for Verifies if this VNET is compatible with an App Service Environment by analyzing the Network Security Group rules.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/VerifyHostingEnvironmentVnet.json
 */
async function verifyHostingEnvironmentVnet() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const parameters = {
    vnetName: "vNet123",
    vnetResourceGroup: "vNet123rg",
    vnetSubnetName: "vNet123SubNet",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.verifyHostingEnvironmentVnet(parameters);
  console.log(result);
}
