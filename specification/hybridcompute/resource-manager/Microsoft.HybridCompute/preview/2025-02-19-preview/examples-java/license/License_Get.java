
/**
 * Samples for Licenses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/license/
     * License_Get.json
     */
    /**
     * Sample code: Get License.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getLicense(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenses().getByResourceGroupWithResponse("myResourceGroup", "{licenseName}",
            com.azure.core.util.Context.NONE);
    }
}
