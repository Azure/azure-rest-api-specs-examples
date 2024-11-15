
/**
 * Samples for LicenseProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/
     * licenseProfile/LicenseProfile_Get.json
     */
    /**
     * Sample code: Get License Profile.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void getLicenseProfile(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenseProfiles().getWithResponse("myResourceGroup", "myMachine", com.azure.core.util.Context.NONE);
    }
}
