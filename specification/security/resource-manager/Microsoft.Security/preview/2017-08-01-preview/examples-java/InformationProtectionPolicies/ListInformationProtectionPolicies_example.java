
/**
 * Samples for InformationProtectionPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/
     * InformationProtectionPolicies/ListInformationProtectionPolicies_example.json
     */
    /**
     * Sample code: Get information protection policies.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getInformationProtectionPolicies(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.informationProtectionPolicies().list(
            "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e",
            com.azure.core.util.Context.NONE);
    }
}
