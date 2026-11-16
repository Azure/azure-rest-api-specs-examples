
/**
 * Samples for FabricShortcuts Approve.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutApprove.json
     */
    /**
     * Sample code: Approve a Fabric shortcut.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void approveAFabricShortcut(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getFabricShortcuts().approveWithResponse("contoso-rg", "contoso-eventhub", "orders",
            "orders-shortcut", com.azure.core.util.Context.NONE);
    }
}
