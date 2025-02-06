
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-22-preview/Organizations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to PineconeVectorDbManager.
     */
    public static void organizationsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.pineconevectordb.PineconeVectorDbManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
