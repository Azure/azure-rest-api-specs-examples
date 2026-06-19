
/**
 * Samples for RoutePolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/RoutePolicies_ListByResourceGroup.json
     */
    /**
     * Sample code: RoutePolicies_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void routePoliciesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.routePolicies().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
