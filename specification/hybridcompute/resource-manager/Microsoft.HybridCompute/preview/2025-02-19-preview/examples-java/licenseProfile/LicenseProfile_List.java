
/**
 * Samples for LicenseProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * licenseProfile/LicenseProfile_List.json
     */
    /**
     * Sample code: List all License Profiles.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void listAllLicenseProfiles(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenseProfiles().list("myResourceGroup", "myMachine", com.azure.core.util.Context.NONE);
    }
}
