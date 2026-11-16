
/**
 * Samples for FabricShortcuts Reject.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutReject.json
     */
    /**
     * Sample code: Reject a Fabric shortcut.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void rejectAFabricShortcut(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getFabricShortcuts().rejectWithResponse("contoso-rg", "contoso-eventhub", "orders",
            "orders-shortcut", com.azure.core.util.Context.NONE);
    }
}
