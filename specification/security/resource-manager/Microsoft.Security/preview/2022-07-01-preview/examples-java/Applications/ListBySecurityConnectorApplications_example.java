
/**
 * Samples for SecurityConnectorApplications List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/
     * ListBySecurityConnectorApplications_example.json
     */
    /**
     * Sample code: List security applications by security connector level scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityApplicationsBySecurityConnectorLevelScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityConnectorApplications().list("gcpResourceGroup", "gcpconnector",
            com.azure.core.util.Context.NONE);
    }
}
