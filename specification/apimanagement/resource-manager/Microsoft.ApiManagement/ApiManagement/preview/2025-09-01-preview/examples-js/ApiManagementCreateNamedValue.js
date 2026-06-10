const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates named value.
 *
 * @summary creates or updates named value.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateNamedValue.json
 */
async function apiManagementCreateNamedValue() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.namedValue.createOrUpdate("rg1", "apimService1", "testprop2", {
    displayName: "prop3name",
    secret: false,
    tags: ["foo", "bar"],
    value: "propValue",
  });
  console.log(result);
}
