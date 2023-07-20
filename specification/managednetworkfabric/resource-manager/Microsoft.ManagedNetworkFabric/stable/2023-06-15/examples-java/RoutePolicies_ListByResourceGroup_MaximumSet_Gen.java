/** Samples for RoutePolicies ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_ListByResourceGroup_MaximumSet_Gen.json
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
