const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to rejects a Microsoft Fabric shortcut.
 *
 * @summary rejects a Microsoft Fabric shortcut.
 * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutReject.json
 */
async function rejectAFabricShortcut() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.fabricShortcuts.reject(
    "contoso-rg",
    "contoso-eventhub",
    "orders",
    "orders-shortcut",
  );
  console.log(result);
}
