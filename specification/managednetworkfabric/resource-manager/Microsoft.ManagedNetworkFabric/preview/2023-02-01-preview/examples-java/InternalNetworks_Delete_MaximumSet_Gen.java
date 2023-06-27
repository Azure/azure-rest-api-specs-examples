/** Samples for InternalNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internalNetworks()
            .delete(
                "resourceGroupName", "example-l3domain", "example-internalnetwork", com.azure.core.util.Context.NONE);
    }
}
