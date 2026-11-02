
/**
 * Samples for Profiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Profiles_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Profiles_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profilesListBySubscriptionMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profiles().list(com.azure.core.util.Context.NONE);
    }
}
