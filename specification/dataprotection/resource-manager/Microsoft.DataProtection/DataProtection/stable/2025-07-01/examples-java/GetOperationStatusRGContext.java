
/**
 * Samples for OperationStatusResourceGroupContext GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/GetOperationStatusRGContext.json
     */
    /**
     * Sample code: Get OperationStatus.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getOperationStatus(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.operationStatusResourceGroupContexts().getByResourceGroupWithResponse("SampleResourceGroup",
            "MjkxOTMyODMtYTE3My00YzJjLTg5NjctN2E4MDIxNDA3NjA2OzdjNGE2ZWRjLWJjMmItNDRkYi1hYzMzLWY1YzEwNzk5Y2EyOA==",
            com.azure.core.util.Context.NONE);
    }
}
