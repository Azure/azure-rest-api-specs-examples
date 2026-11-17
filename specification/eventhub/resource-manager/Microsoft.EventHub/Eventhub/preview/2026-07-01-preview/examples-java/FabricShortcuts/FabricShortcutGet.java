
/**
 * Samples for FabricShortcuts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutGet.json
     */
    /**
     * Sample code: Get a Fabric shortcut.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void getAFabricShortcut(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getFabricShortcuts().getWithResponse("contoso-rg", "contoso-eventhub", "orders",
            "orders-shortcut", com.azure.core.util.Context.NONE);
    }
}
