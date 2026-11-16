
/**
 * Samples for ApplicationGroup Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ApplicationGroup/ApplicationGroupDelete.json
     */
    /**
     * Sample code: ApplicationGroupDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void applicationGroupDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getApplicationGroups().deleteWithResponse("contosotest", "contoso-ua-test-eh-system-1",
            "appGroup1", com.azure.core.util.Context.NONE);
    }
}
