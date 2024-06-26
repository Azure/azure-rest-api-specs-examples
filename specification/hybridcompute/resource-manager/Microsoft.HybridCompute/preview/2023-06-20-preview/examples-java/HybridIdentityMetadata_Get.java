/** Samples for HybridIdentityMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/HybridIdentityMetadata_Get.json
     */
    /**
     * Sample code: GetHybridIdentityMetadata.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getHybridIdentityMetadata(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .hybridIdentityMetadatas()
            .getWithResponse("testrg", "ContosoVm", "default", com.azure.core.util.Context.NONE);
    }
}
