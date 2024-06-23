const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Approves or rejects a private endpoint connection
 *
 * @summary Description for Approves or rejects a private endpoint connection
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ApproveRejectSitePrivateEndpointConnection.json
 */
async function approvesOrRejectsAPrivateEndpointConnectionForASite() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testSite";
  const privateEndpointConnectionName = "connection";
  const privateEndpointWrapper = {
    privateLinkServiceConnectionState: {
      description: "Approved by admin.",
      actionsRequired: "",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.beginApproveOrRejectPrivateEndpointConnectionAndWait(
    resourceGroupName,
    name,
    privateEndpointConnectionName,
    privateEndpointWrapper,
  );
  console.log(result);
}
