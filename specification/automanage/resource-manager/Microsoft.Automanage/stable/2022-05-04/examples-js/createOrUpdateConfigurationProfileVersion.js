const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a configuration profile version
 *
 * @summary Creates a configuration profile version
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileVersion.json
 */
async function createOrUpdateConfigurationProfileVersion() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const versionName = "version1";
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
  const result = await client.configurationProfilesVersions.createOrUpdate(
    configurationProfileName,
    versionName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

createOrUpdateConfigurationProfileVersion().catch(console.error);
