/** Samples for AccessControlLists List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/AccessControlLists_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: AccessControlLists_ListBySubscription_MinimumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsListBySubscriptionMinimumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().list(com.azure.core.util.Context.NONE);
    }
}
