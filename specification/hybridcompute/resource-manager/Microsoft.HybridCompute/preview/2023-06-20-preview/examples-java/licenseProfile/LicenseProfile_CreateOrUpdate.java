/** Samples for LicenseProfiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/licenseProfile/LicenseProfile_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update a License Profile.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void createOrUpdateALicenseProfile(
        com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager
            .licenseProfiles()
            .define()
            .withRegion("eastus2euap")
            .withExistingMachine("myResourceGroup", "myMachine")
            .withAssignedLicense("{LicenseResourceId}")
            .create();
    }
}
