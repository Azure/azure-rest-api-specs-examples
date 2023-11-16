/** Samples for HybridIdentityMetadata ListByMachines. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/HybridIdentityMetadata_ListByVirtualMachines.json
     */
    /**
     * Sample code: HybridIdentityMetadataListByVirtualMachines.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void hybridIdentityMetadataListByVirtualMachines(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.hybridIdentityMetadatas().listByMachines("testrg", "ContosoVm", com.azure.core.util.Context.NONE);
    }
}
