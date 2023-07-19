/** Samples for ExternalNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_Get_MaximumSet_Gen.json
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
            .getWithResponse("rgL3IsolationDomains", "yhtr", "fltpszzikbalrzaqq", com.azure.core.util.Context.NONE);
    }
}
