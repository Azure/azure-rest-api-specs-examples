
import java.time.Duration;

/**
 * Samples for Metrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetricError.json
     */
    /**
     * Sample code: Get Metric with error.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricWithError(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listWithResponse(
            "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo",
            "2021-06-07T21:51:00Z/2021-06-08T01:51:00Z", Duration.parse("FULL"), "MongoRequestsCount,MongoRequests",
            "average", null, null, null, null, "microsoft.documentdb/databaseaccounts", true, false, null,
            com.azure.core.util.Context.NONE);
    }
}
