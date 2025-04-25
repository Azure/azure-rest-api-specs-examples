const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates named value.
 *
 * @summary Creates or updates named value.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceNamedValue.json
 */
async function apiManagementCreateWorkspaceNamedValue() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const namedValueId = "testprop2";
  const parameters = {
    displayName: "prop3name",
    secret: false,
    tags: ["foo", "bar"],
    value: "propValue",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceNamedValue.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    workspaceId,
    namedValueId,
    parameters,
  );
  console.log(result);
}
