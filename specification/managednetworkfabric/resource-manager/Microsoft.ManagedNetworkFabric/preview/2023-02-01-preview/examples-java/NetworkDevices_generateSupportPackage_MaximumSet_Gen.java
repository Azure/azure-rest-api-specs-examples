/** Samples for NetworkDevices GenerateSupportPackage. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_generateSupportPackage_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkDevices_generateSupportPackage_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkDevicesGenerateSupportPackageMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkDevices()
            .generateSupportPackage("resourceGroupName", "networkDeviceName", com.azure.core.util.Context.NONE);
    }
}
