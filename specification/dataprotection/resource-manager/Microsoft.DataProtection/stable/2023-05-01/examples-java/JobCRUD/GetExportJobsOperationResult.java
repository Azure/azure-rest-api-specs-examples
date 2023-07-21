/** Samples for ExportJobsOperationResult Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/JobCRUD/GetExportJobsOperationResult.json
     */
    /**
     * Sample code: Get Export Jobs Operation Result.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getExportJobsOperationResult(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .exportJobsOperationResults()
            .getWithResponse(
                "SwaggerTestRg",
                "NetSDKTestRsVault",
                "00000000-0000-0000-0000-000000000000",
                com.azure.core.util.Context.NONE);
    }
}
