/** Samples for NetworkFabricControllers EnableWorkloadManagementNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_enableWorkloadManagementNetwork_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_enableWorkloadManagementNetwork_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersEnableWorkloadManagementNetworkMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabricControllers()
            .enableWorkloadManagementNetwork(
                "resourceGroupName", "networkFabricControllerName", com.azure.core.util.Context.NONE);
    }
}
