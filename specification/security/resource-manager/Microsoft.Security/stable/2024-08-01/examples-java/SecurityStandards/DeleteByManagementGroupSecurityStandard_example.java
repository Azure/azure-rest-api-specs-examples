
/**
 * Samples for SecurityStandards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * DeleteByManagementGroupSecurityStandard_example.json
     */
    /**
     * Sample code: Delete a security standard over management group scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        deleteASecurityStandardOverManagementGroupScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().deleteByResourceGroupWithResponse(
            "providers/Microsoft.Management/managementGroups/contoso", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
            com.azure.core.util.Context.NONE);
    }
}
