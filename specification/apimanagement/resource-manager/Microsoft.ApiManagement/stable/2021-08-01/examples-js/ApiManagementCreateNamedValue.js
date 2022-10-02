const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates named value.
 *
 * @summary Creates or updates named value.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateNamedValue.json
 */
async function apiManagementCreateNamedValue() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const namedValueId = "testprop2";
  const parameters = {
    displayName: "prop3name",
    secret: false,
    tags: ["foo", "bar"],
    value: "propValue",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.namedValue.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    namedValueId,
    parameters
  );
  console.log(result);
}

apiManagementCreateNamedValue().catch(console.error);
