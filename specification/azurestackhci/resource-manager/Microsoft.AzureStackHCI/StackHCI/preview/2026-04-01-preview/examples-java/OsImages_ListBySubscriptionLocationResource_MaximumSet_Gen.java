
/**
 * Samples for OsImages ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/OsImages_ListBySubscriptionLocationResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: OsImages_ListBySubscriptionLocationResource_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void osImagesListBySubscriptionLocationResourceMaximumSet(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.osImages().listBySubscriptionLocationResource("westus2", com.azure.core.util.Context.NONE);
    }
}
