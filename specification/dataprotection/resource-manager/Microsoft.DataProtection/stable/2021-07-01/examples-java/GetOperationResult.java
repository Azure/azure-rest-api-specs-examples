import com.azure.core.util.Context;

/** Samples for OperationResult Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/GetOperationResult.json
     */
    /**
     * Sample code: Get OperationResult.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getOperationResult(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .operationResults()
            .getWithResponse(
                "MjkxOTMyODMtYTE3My00YzJjLTg5NjctN2E4MDIxNDA3NjA2OzdjNGE2ZWRjLWJjMmItNDRkYi1hYzMzLWY1YzEwNzk5Y2EyOA==",
                "WestUS",
                Context.NONE);
    }
}
