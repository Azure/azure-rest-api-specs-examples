/** Samples for InternalNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internalNetworks()
            .getWithResponse(
                "example-rg", "example-l3domain", "example-internalnetwork", com.azure.core.util.Context.NONE);
    }
}
