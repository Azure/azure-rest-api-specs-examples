const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a configuration profile
 *
 * @summary Updates a configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/updateConfigurationProfile.json
 */
async function updateConfigurationProfile() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const resourceGroupName = "myResourceGroupName";
  const parameters = {
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
  const result = await client.configurationProfiles.update(
    configurationProfileName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

updateConfigurationProfile().catch(console.error);
