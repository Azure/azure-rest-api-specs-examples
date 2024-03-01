const { AzureAPICenter } = require("@azure/arm-apicenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates existing service.
 *
 * @summary Updates existing service.
 * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Services_Update.json
 */
async function servicesUpdate() {
  const subscriptionId =
    process.env["APICENTER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APICENTER_RESOURCE_GROUP"] || "contoso-resources";
  const serviceName = "contoso";
  const payload = {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/contosoResources/providers/MicrosoftManagedIdentity/userAssignedIdentities/contosoIdentity":
          {},
      },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureAPICenter(credential, subscriptionId);
  const result = await client.services.update(resourceGroupName, serviceName, payload);
  console.log(result);
}
