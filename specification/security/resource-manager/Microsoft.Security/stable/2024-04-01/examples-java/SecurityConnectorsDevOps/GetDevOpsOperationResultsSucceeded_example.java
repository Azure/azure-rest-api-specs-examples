
/**
 * Samples for DevOpsOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * GetDevOpsOperationResultsSucceeded_example.json
     */
    /**
     * Sample code: Get_DevOpsOperationResults_Succeeded.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDevOpsOperationResultsSucceeded(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsOperationResults().getWithResponse("myRg", "mySecurityConnectorName",
            "4e826cf1-5c36-4808-a7d2-fb4f5170978b", com.azure.core.util.Context.NONE);
    }
}
