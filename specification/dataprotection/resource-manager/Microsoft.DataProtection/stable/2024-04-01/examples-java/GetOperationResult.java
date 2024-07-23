
/**
 * Samples for OperationResult Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * GetOperationResult.json
     */
    /**
     * Sample code: Get OperationResult.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.operationResults().getWithResponse(
            "MjkxOTMyODMtYTE3My00YzJjLTg5NjctN2E4MDIxNDA3NjA2OzdjNGE2ZWRjLWJjMmItNDRkYi1hYzMzLWY1YzEwNzk5Y2EyOA==",
            "WestUS", com.azure.core.util.Context.NONE);
    }
}
