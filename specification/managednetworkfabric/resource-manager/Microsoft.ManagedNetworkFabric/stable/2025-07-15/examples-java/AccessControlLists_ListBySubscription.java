
/**
 * Samples for AccessControlLists List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/AccessControlLists_ListBySubscription.json
     */
    /**
     * Sample code: AccessControlLists_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().list(com.azure.core.util.Context.NONE);
    }
}
