
/**
 * Samples for FabricShortcuts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutDelete.json
     */
    /**
     * Sample code: Delete a Fabric shortcut.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void deleteAFabricShortcut(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getFabricShortcuts().deleteWithResponse("contoso-rg", "contoso-eventhub", "orders",
            "orders-shortcut", com.azure.core.util.Context.NONE);
    }
}
