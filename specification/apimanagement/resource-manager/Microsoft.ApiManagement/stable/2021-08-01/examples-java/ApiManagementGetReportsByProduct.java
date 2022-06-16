import com.azure.core.util.Context;

/** Samples for Reports ListByProduct. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetReportsByProduct.json
     */
    /**
     * Sample code: ApiManagementGetReportsByProduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetReportsByProduct(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .reports()
            .listByProduct(
                "rg1",
                "apimService1",
                "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'",
                null,
                null,
                null,
                Context.NONE);
    }
}
