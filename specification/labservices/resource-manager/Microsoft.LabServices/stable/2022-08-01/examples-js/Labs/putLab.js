const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to create or update a lab resource.
 *
 * @summary Operation to create or update a lab resource.
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/putLab.json
 */
async function putLab() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const labName = "testlab";
  const body = {
    description: "This is a test lab.",
    autoShutdownProfile: {
      disconnectDelay: "PT5M",
      idleDelay: "PT5M",
      noConnectDelay: "PT5M",
      shutdownOnDisconnect: "Enabled",
      shutdownOnIdle: "UserAbsence",
      shutdownWhenNotConnected: "Enabled",
    },
    connectionProfile: {
      clientRdpAccess: "Public",
      clientSshAccess: "Public",
      webRdpAccess: "None",
      webSshAccess: "None",
    },
    labPlanId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan",
    location: "westus",
    networkProfile: {
      subnetId:
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
    },
    securityProfile: { openAccess: "Disabled" },
    state: "Draft",
    title: "Test Lab",
    virtualMachineProfile: {
      additionalCapabilities: { installGpuDrivers: "Disabled" },
      adminUser: { username: "test-user" },
      createOption: "TemplateVM",
      imageReference: {
        offer: "WindowsServer",
        publisher: "Microsoft",
        sku: "2019-Datacenter",
        version: "2019.0.20190410",
      },
      sku: { name: "Medium" },
      usageQuota: "PT10H",
      useSharedPassword: "Disabled",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const result = await client.labs.beginCreateOrUpdateAndWait(resourceGroupName, labName, body);
  console.log(result);
}

putLab().catch(console.error);
