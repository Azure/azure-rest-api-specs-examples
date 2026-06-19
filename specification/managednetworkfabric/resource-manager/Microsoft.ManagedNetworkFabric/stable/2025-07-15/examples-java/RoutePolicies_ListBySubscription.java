
/**
 * Samples for RoutePolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/RoutePolicies_ListBySubscription.json
     */
    /**
     * Sample code: RoutePolicies_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void routePoliciesListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.routePolicies().list(com.azure.core.util.Context.NONE);
    }
}
