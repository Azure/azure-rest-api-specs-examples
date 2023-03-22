import com.azure.resourcemanager.security.models.InformationProtectionPolicyName;

/** Samples for InformationProtectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/GetCustomInformationProtectionPolicy_example.json
     */
    /**
     * Sample code: Get the customized information protection policy for a management group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getTheCustomizedInformationProtectionPolicyForAManagementGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .informationProtectionPolicies()
            .getWithResponse(
                "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e",
                InformationProtectionPolicyName.CUSTOM,
                com.azure.core.util.Context.NONE);
    }
}
