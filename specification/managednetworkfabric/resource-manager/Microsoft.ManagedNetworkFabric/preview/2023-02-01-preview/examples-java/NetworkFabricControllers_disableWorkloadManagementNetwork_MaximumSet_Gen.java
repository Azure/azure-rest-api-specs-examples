/** Samples for NetworkFabricControllers DisableWorkloadManagementNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_disableWorkloadManagementNetwork_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_disableWorkloadManagementNetwork_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersDisableWorkloadManagementNetworkMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabricControllers()
            .disableWorkloadManagementNetwork(
                "resourceGroupName", "networkFabricControllerName", com.azure.core.util.Context.NONE);
    }
}
