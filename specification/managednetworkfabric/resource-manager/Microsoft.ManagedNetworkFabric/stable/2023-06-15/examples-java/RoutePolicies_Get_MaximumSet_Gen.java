/** Samples for RoutePolicies GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: RoutePolicies_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void routePoliciesGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .routePolicies()
            .getByResourceGroupWithResponse("example-rg", "example-routePolicy", com.azure.core.util.Context.NONE);
    }
}
