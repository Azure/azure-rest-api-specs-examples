
/**
 * Samples for Organization List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListBySubscriptionMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
