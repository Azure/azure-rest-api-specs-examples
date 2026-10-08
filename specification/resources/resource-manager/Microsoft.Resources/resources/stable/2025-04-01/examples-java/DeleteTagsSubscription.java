
/**
 * Samples for TagOperations DeleteAtScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/DeleteTagsSubscription.json
     */
    /**
     * Sample code: Update tags on a subscription.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void updateTagsOnASubscription(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getTagOperations().deleteAtScope("subscriptions/00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
