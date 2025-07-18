const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new storage appliance or update the properties of the existing one.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 *
 * @summary Create a new storage appliance or update the properties of the existing one.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/StorageAppliances_Create.json
 */
async function createOrUpdateStorageAppliance() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const storageApplianceName = "storageApplianceName";
  const storageApplianceParameters = {
    administratorCredentials: { password: "{password}", username: "adminUser" },
    extendedLocation: {
      name: "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
      type: "CustomLocation",
    },
    location: "location",
    rackId:
      "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName",
    rackSlot: 1,
    serialNumber: "BM1219XXX",
    storageApplianceSkuId: "684E-3B16-399E",
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.storageAppliances.beginCreateOrUpdateAndWait(
    resourceGroupName,
    storageApplianceName,
    storageApplianceParameters,
  );
  console.log(result);
}
