/** Samples for Reports ListByRequest. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetReportsByRequest.json
     */
    /**
     * Sample code: ApiManagementGetReportsByRequest.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetReportsByRequest(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .reports()
            .listByRequest(
                "rg1",
                "apimService1",
                "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
