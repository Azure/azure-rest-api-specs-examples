const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an integration account map. If the map is larger than 4 MB, you need to store the map in an Azure blob and use the blob's Shared Access Signature (SAS) URL as the 'contentLink' property value.
 *
 * @summary Creates or updates an integration account map. If the map is larger than 4 MB, you need to store the map in an Azure blob and use the blob's Shared Access Signature (SAS) URL as the 'contentLink' property value.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountLargeMaps_CreateOrUpdate.json
 */
async function createOrUpdateAMapLargerThan4Mb() {
  const subscriptionId = process.env["LOGIC_SUBSCRIPTION_ID"] || "<Azure-subscription-ID>";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const integrationAccountName = "testIntegrationAccount";
  const mapName = "testMap";
  const map = {
    contentLink: { uri: "<blob-SAS-URL-for-map>" },
    contentType: "application/xml",
    location: "westus",
    mapType: "Xslt",
    metadata: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationAccountMaps.createOrUpdate(
    resourceGroupName,
    integrationAccountName,
    mapName,
    map
  );
  console.log(result);
}
