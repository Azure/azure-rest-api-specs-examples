
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PrivateLinkResources_Get.
     * json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.privateLinkResources().getWithResponse("examplerg", "topics", "exampletopic1", "topic",
            com.azure.core.util.Context.NONE);
    }
}
