import com.azure.core.util.Context;

/** Samples for SecurityPolicies ListByProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_ListByProfile.json
     */
    /**
     * Sample code: SecurityPolicies_ListByProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void securityPoliciesListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getSecurityPolicies()
            .listByProfile("RG", "profile1", Context.NONE);
    }
}
