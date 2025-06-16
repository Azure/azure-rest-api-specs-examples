
/**
 * Samples for LicenseProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * licenseProfile/LicenseProfile_Delete.json
     */
    /**
     * Sample code: Delete a License Profile.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void deleteALicenseProfile(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenseProfiles().delete("myResourceGroup", "myMachine", com.azure.core.util.Context.NONE);
    }
}
