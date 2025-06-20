
import com.azure.resourcemanager.hybridcompute.models.LicenseProfileProductType;
import com.azure.resourcemanager.hybridcompute.models.LicenseProfileSubscriptionStatus;
import com.azure.resourcemanager.hybridcompute.models.ProductFeature;
import java.util.Arrays;

/**
 * Samples for LicenseProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/
     * licenseProfile/LicenseProfile_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update a License Profile.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        createOrUpdateALicenseProfile(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenseProfiles().define().withRegion("eastus2euap").withExistingMachine("myResourceGroup", "myMachine")
            .withSoftwareAssuranceCustomer(true).withAssignedLicense("{LicenseResourceId}")
            .withSubscriptionStatus(LicenseProfileSubscriptionStatus.ENABLED)
            .withProductType(LicenseProfileProductType.WINDOWS_SERVER)
            .withProductFeatures(Arrays.asList(new ProductFeature().withName("Hotpatch")
                .withSubscriptionStatus(LicenseProfileSubscriptionStatus.ENABLED)))
            .create();
    }
}
