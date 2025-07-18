const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new cloud services network or update the properties of the existing cloud services network.
 *
 * @summary Create a new cloud services network or update the properties of the existing cloud services network.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/CloudServicesNetworks_Create.json
 */
async function createOrUpdateCloudServicesNetwork() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const cloudServicesNetworkName = "cloudServicesNetworkName";
  const cloudServicesNetworkParameters = {
    additionalEgressEndpoints: [
      {
        category: "azure-resource-management",
        endpoints: [{ domainName: "storageaccountex.blob.core.windows.net", port: 443 }],
      },
    ],
    enableDefaultEgressEndpoints: "False",
    extendedLocation: {
      name: "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
      type: "CustomLocation",
    },
    location: "location",
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.cloudServicesNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    cloudServicesNetworkName,
    cloudServicesNetworkParameters,
  );
  console.log(result);
}
