
/**
 * Samples for Licenses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/license/
     * License_Delete.json
     */
    /**
     * Sample code: Delete a License.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteALicense(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenses().delete("myResourceGroup", "{licenseName}", com.azure.core.util.Context.NONE);
    }
}
