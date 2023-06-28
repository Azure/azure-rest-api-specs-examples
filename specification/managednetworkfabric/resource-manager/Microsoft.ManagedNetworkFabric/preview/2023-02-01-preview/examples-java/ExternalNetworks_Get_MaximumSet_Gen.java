/** Samples for ExternalNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalNetworks_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .externalNetworks()
            .getWithResponse(
                "resourceGroupName", "example-l3domain", "example-externalnetwork", com.azure.core.util.Context.NONE);
    }
}
