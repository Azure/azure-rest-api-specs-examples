
/**
 * Samples for SecurityStandards List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * ListByManagementGroupSecurityStandards_example.json
     */
    /**
     * Sample code: List security standards by management group scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        listSecurityStandardsByManagementGroupScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().list("providers/Microsoft.Management/managementGroups/contoso",
            com.azure.core.util.Context.NONE);
    }
}
