const { AtlasClient } = require("@azure/arm-mongodbatlas");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a OrganizationResource
 *
 * @summary delete a OrganizationResource
 * x-ms-original-file: 2024-11-18-preview/Organizations_Delete_MaximumSet_Gen.json
 */
async function organizationsDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4AFC1287-D389-4265-B2D4-59B96A45CACC";
  const client = new AtlasClient(credential, subscriptionId);
  await client.organizations.delete("rgopenapi", "U.1-:7");
}
