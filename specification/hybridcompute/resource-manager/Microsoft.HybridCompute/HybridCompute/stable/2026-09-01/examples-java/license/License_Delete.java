
/**
 * Samples for Licenses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/license/License_Delete.json
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
