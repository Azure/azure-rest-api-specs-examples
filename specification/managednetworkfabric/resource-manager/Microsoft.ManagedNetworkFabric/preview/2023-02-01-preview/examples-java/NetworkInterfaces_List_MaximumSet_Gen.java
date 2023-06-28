/** Samples for NetworkInterfaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkInterfaces_List_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkInterfacesListMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkInterfaces().list("resourceGroupName", "networkDeviceName", com.azure.core.util.Context.NONE);
    }
}
