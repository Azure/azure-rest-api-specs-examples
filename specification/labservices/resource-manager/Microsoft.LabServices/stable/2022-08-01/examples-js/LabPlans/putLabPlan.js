const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to create or update a Lab Plan resource.
 *
 * @summary Operation to create or update a Lab Plan resource.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/putLabPlan.json
 */
async function putLabPlan() {
  const subscriptionId =
    process.env["LABSERVICES_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LABSERVICES_RESOURCE_GROUP"] || "testrg123";
  const labPlanName = "testlabplan";
  const body = {
    defaultAutoShutdownProfile: {
      disconnectDelay: "PT5M",
      idleDelay: "PT5M",
      noConnectDelay: "PT5M",
      shutdownOnDisconnect: "Enabled",
      shutdownOnIdle: "UserAbsence",
      shutdownWhenNotConnected: "Enabled",
    },
    defaultConnectionProfile: {
      clientRdpAccess: "Public",
      clientSshAccess: "Public",
      webRdpAccess: "None",
      webSshAccess: "None",
    },
    defaultNetworkProfile: {
      subnetId:
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
    },
    location: "westus",
    sharedGalleryId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig",
    supportInfo: {
      email: "help@contoso.com",
      instructions: "Contact support for help.",
      phone: "+1-202-555-0123",
      url: "help.contoso.com",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labPlans.beginCreateOrUpdateAndWait(
    resourceGroupName,
    labPlanName,
    body
  );
  console.log(result);
}
