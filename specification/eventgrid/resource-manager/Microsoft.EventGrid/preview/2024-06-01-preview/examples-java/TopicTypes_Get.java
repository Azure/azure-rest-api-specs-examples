
/**
 * Samples for TopicTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/TopicTypes_Get.
     * json
     */
    /**
     * Sample code: TopicTypes_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().getWithResponse("Microsoft.Storage.StorageAccounts", com.azure.core.util.Context.NONE);
    }
}
