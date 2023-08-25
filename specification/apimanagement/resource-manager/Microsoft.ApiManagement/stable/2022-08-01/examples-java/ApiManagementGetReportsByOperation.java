/** Samples for Reports ListByOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetReportsByOperation.json
     */
    /**
     * Sample code: ApiManagementGetReportsByOperation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetReportsByOperation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .reports()
            .listByOperation(
                "rg1",
                "apimService1",
                "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'",
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
