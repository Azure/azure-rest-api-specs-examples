const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a configuration profile
 *
 * @summary Creates a configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfile.json
 */
async function createOrUpdateConfigurationProfile() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const resourceGroupName = "myResourceGroupName";
  const parameters = {
    location: "East US",
    properties: {
      configuration: {
        "Antimalware/Enable": false,
        "AzureSecurityCenter/Enable": true,
        "Backup/Enable": false,
        "BootDiagnostics/Enable": true,
        "ChangeTrackingAndInventory/Enable": true,
        "GuestConfiguration/Enable": true,
        "LogAnalytics/Enable": true,
        "UpdateManagement/Enable": true,
        "VMInsights/Enable": true,
      },
    },
    tags: { organization: "Administration" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfiles.createOrUpdate(
    configurationProfileName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

createOrUpdateConfigurationProfile().catch(console.error);
