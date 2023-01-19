/** Samples for DiskPools ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_GetOutboundNetworkDependencies.json
     */
    /**
     * Sample code: Get Disk Pool outbound network dependencies.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void getDiskPoolOutboundNetworkDependencies(
        com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager
            .diskPools()
            .listOutboundNetworkDependenciesEndpoints(
                "Sample-WestUSResourceGroup", "SampleAse", com.azure.core.util.Context.NONE);
    }
}
