
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: NameSpacePrivateLinkResourcesGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpacePrivateLinkResourcesGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getPrivateLinkResources().getWithResponse("ArunMonocle", "sdk-Namespace-2924",
            com.azure.core.util.Context.NONE);
    }
}
