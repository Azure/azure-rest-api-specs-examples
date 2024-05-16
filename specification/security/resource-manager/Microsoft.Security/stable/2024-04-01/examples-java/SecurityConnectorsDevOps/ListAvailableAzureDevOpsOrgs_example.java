
/**
 * Samples for AzureDevOpsOrgs ListAvailable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListAvailableAzureDevOpsOrgs_example.json
     */
    /**
     * Sample code: ListAvailable_AzureDevOpsOrgs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAvailableAzureDevOpsOrgs(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsOrgs().listAvailableWithResponse("myRg", "mySecurityConnectorName",
            com.azure.core.util.Context.NONE);
    }
}
