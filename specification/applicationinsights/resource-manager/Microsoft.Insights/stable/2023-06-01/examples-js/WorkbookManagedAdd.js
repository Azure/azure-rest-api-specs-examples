const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new workbook.
 *
 * @summary Create a new workbook.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2023-06-01/examples/WorkbookManagedAdd.json
 */
async function workbookManagedAdd() {
  const subscriptionId =
    process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const sourceId =
    "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group";
  const workbookProperties = {
    description: "Sample workbook",
    category: "workbook",
    displayName: "Sample workbook",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/myid":
          {},
      },
    },
    kind: "shared",
    location: "westus",
    serializedData:
      '{"version":"Notebook/1.0","items":[{"type":1,"content":{"json":"test"},"name":"text - 0"}],"isLocked":false,"fallbackResourceIds":["/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/my-resource-group"]}',
    storageUri:
      "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorage/blobServices/default/containers/mycontainer",
    version: "Notebook/1.0",
  };
  const options = { sourceId };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbooks.createOrUpdate(
    resourceGroupName,
    resourceName,
    workbookProperties,
    options,
  );
  console.log(result);
}
