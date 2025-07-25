const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Patch properties of the provided rack, or update the tags associated with the rack. Properties and tag updates can be done independently.
 *
 * @summary Patch properties of the provided rack, or update the tags associated with the rack. Properties and tag updates can be done independently.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Racks_Patch.json
 */
async function patchRack() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const rackName = "rackName";
  const rackUpdateParameters = {
    rackLocation: "Rack 2B",
    rackSerialNumber: "RACK_SERIAL_NUMBER",
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const options = { rackUpdateParameters };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.racks.beginUpdateAndWait(resourceGroupName, rackName, options);
  console.log(result);
}
