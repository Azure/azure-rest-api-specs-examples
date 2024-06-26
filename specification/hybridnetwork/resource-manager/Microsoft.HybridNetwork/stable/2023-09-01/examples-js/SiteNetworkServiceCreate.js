const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network site.
 *
 * @summary Creates or updates a network site.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteNetworkServiceCreate.json
 */
async function createSiteNetworkService() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg1";
  const siteNetworkServiceName = "testSiteNetworkServiceName";
  const parameters = {
    location: "westUs2",
    properties: {
      desiredStateConfigurationGroupValueReferences: {
        myVMConfiguration: {
          id: "/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/configurationgroupvalues/MyVM_Configuration1",
        },
      },
      networkServiceDesignVersionResourceReference: {
        id: "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkServiceDesignGroups/TestNetworkServiceDesignGroupName/networkServiceDesignVersions/1.0.0",
        idType: "Open",
      },
      siteReference: {
        id: "/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/sites/testSite",
      },
    },
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.siteNetworkServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    siteNetworkServiceName,
    parameters,
  );
  console.log(result);
}
