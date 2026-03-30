
/**
 * Samples for ResourceHealthMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListResourceHealthMetadataBySubscription.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listResourceHealthMetadataForASubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().list(com.azure.core.util.Context.NONE);
    }
}
