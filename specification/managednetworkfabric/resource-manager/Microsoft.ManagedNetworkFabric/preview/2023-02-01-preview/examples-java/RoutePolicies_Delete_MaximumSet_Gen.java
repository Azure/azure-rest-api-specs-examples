/** Samples for RoutePolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/RoutePolicies_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RoutePolicies_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void routePoliciesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.routePolicies().delete("rgRoutePolicies", "routePolicyName", com.azure.core.util.Context.NONE);
    }
}
