const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function upgradeMachineExtensions() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const extensionUpgradeParameters = {
    extensionTargets: {
      microsoftAzureMonitoring: { targetVersion: "2.0" },
      microsoftComputeCustomScriptExtension: { targetVersion: "1.10" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.beginUpgradeExtensionsAndWait(
    resourceGroupName,
    machineName,
    extensionUpgradeParameters
  );
  console.log(result);
}

upgradeMachineExtensions().catch(console.error);
