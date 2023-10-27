/** Samples for SecurityPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/SecurityPolicies_Delete.json
     */
    /**
     * Sample code: SecurityPolicies_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void securityPoliciesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getSecurityPolicies()
            .delete("RG", "profile1", "securityPolicy1", com.azure.core.util.Context.NONE);
    }
}
