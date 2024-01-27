
/** Samples for SecurityPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/SecurityPolicies_Get.json
     */
    /**
     * Sample code: SecurityPolicies_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void securityPoliciesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getSecurityPolicies().getWithResponse("RG", "profile1",
            "securityPolicy1", com.azure.core.util.Context.NONE);
    }
}
