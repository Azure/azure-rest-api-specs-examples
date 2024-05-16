
/**
 * Samples for DevOpsOperationResults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * GetDevOpsOperationResultsFailed_example.json
     */
    /**
     * Sample code: Get_DevOpsOperationResults_Failed.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDevOpsOperationResultsFailed(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsOperationResults().getWithResponse("myRg", "mySecurityConnectorName",
            "8d4caace-e7b3-4b3e-af99-73f76829ebcf", com.azure.core.util.Context.NONE);
    }
}
