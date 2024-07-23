
/**
 * Samples for OperationStatusResourceGroupContext GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * GetOperationStatusRGContext.json
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
