const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns ResourceGuardProxy under vault and with the name referenced in request
 *
 * @summary Returns ResourceGuardProxy under vault and with the name referenced in request
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
 */
async function getResourceGuardProxy() {
  const subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const vaultName = "sampleVault";
  const resourceGroupName = "SampleResourceGroup";
  const resourceGuardProxyName = "swaggerExample";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.resourceGuardProxy.get(
    vaultName,
    resourceGroupName,
    resourceGuardProxyName
  );
  console.log(result);
}

getResourceGuardProxy().catch(console.error);
