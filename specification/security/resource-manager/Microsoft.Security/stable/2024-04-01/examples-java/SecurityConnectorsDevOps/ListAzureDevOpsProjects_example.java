
/**
 * Samples for AzureDevOpsProjects List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListAzureDevOpsProjects_example.json
     */
    /**
     * Sample code: List_AzureDevOpsProjects.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAzureDevOpsProjects(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsProjects().list("myRg", "mySecurityConnectorName", "myAzDevOpsOrg",
            com.azure.core.util.Context.NONE);
    }
}
