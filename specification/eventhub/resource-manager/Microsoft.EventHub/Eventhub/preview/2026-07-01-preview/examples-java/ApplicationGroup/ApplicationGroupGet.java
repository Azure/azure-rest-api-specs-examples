
/**
 * Samples for ApplicationGroup Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/ApplicationGroup/ApplicationGroupGet.json
     */
    /**
     * Sample code: ApplicationGroupGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void applicationGroupGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getApplicationGroups().getWithResponse("contosotest", "contoso-ua-test-eh-system-1",
            "appGroup1", com.azure.core.util.Context.NONE);
    }
}
