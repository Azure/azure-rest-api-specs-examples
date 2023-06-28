/** Samples for RoutePolicies GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/RoutePolicies_Get_MaximumSet_Gen.json
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
            .getByResourceGroupWithResponse("rgRoutePolicies", "routePolicyName", com.azure.core.util.Context.NONE);
    }
}
