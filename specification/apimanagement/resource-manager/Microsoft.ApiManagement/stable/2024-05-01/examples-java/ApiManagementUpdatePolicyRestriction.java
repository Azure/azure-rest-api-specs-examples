
import com.azure.resourcemanager.apimanagement.models.PolicyRestrictionContract;

/**
 * Samples for PolicyRestriction Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdatePolicyRestriction.json
     */
    /**
     * Sample code: ApiManagementUpdatePolicyRestriction.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdatePolicyRestriction(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        PolicyRestrictionContract resource = manager.policyRestrictions()
            .getWithResponse("rg1", "apimService1", "policyRestriction1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withScope("Sample Path 2 to the policy document.").withIfMatch("*").apply();
    }
}
