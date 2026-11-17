
/**
 * Samples for FabricShortcuts ListByEventHub.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/FabricShortcuts/FabricShortcutList.json
     */
    /**
     * Sample code: List Fabric shortcuts for an Event Hub.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listFabricShortcutsForAnEventHub(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getFabricShortcuts().listByEventHub("contoso-rg", "contoso-eventhub", "orders",
            com.azure.core.util.Context.NONE);
    }
}
