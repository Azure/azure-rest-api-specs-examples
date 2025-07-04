
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Organizations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        organizationsListBySubscriptionMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
