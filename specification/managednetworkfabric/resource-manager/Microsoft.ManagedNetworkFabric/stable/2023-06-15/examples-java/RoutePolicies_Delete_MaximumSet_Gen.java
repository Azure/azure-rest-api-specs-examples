/** Samples for RoutePolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RoutePolicies_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void routePoliciesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.routePolicies().delete("example-rg", "example-routePolicy", com.azure.core.util.Context.NONE);
    }
}
