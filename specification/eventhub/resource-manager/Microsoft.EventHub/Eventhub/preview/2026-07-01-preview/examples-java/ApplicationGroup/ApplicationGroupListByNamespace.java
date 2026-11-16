
/**
 * Samples for ApplicationGroup ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ApplicationGroup/ApplicationGroupListByNamespace.json
     */
    /**
     * Sample code: ListApplicationGroups.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listApplicationGroups(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getApplicationGroups().listByNamespace("contosotest", "contoso-ua-test-eh-system-1",
            com.azure.core.util.Context.NONE);
    }
}
