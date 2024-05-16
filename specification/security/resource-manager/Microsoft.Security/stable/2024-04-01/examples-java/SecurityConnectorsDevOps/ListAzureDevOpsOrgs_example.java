
/**
 * Samples for AzureDevOpsOrgs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListAzureDevOpsOrgs_example.json
     */
    /**
     * Sample code: List_AzureDevOpsOrgs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAzureDevOpsOrgs(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsOrgs().list("myRg", "mySecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
