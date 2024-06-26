import com.azure.resourcemanager.hybridcompute.models.LicenseProfile;

/** Samples for LicenseProfiles Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/licenseProfile/LicenseProfile_Update.json
     */
    /**
     * Sample code: Update a License Profile.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void updateALicenseProfile(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        LicenseProfile resource =
            manager
                .licenseProfiles()
                .getWithResponse("myResourceGroup", "myMachine", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withAssignedLicense("{LicenseResourceId}").apply();
    }
}
