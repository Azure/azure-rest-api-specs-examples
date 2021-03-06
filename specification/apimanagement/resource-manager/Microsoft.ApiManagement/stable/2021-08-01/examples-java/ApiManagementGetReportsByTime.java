import com.azure.core.util.Context;
import java.time.Duration;

/** Samples for Reports ListByTime. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetReportsByTime.json
     */
    /**
     * Sample code: ApiManagementGetReportsByTime.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetReportsByTime(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .reports()
            .listByTime(
                "rg1",
                "apimService1",
                "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'",
                Duration.parse("PT15M"),
                null,
                null,
                null,
                Context.NONE);
    }
}
