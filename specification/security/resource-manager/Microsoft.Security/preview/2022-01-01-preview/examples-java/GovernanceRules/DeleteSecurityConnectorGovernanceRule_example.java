import com.azure.core.util.Context;

/** Samples for SecurityConnectorGovernanceRulesOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/DeleteSecurityConnectorGovernanceRule_example.json
     */
    /**
     * Sample code: Delete security GovernanceRule.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteSecurityGovernanceRule(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securityConnectorGovernanceRulesOperations()
            .deleteWithResponse(
                "gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", Context.NONE);
    }
}
