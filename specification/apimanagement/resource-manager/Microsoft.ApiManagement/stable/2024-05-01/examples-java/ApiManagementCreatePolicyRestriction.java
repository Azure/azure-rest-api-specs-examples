
import com.azure.resourcemanager.apimanagement.models.PolicyRestrictionRequireBase;

/**
 * Samples for PolicyRestriction CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreatePolicyRestriction.json
     */
    /**
     * Sample code: ApiManagementCreatePolicyRestriction.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreatePolicyRestriction(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policyRestrictions().define("policyRestriction1").withExistingService("rg1", "apimService1")
            .withScope("Sample Path to the policy document.").withRequireBase(PolicyRestrictionRequireBase.TRUE)
            .withIfMatch("*").create();
    }
}
