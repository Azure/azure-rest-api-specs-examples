/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Operations_List.json
     */
    /**
     * Sample code: Lists all of the available Data Lake Analytics REST API operations.
     *
     * @param manager Entry point to DataLakeAnalyticsManager.
     */
    public static void listsAllOfTheAvailableDataLakeAnalyticsRESTAPIOperations(
        com.azure.resourcemanager.datalakeanalytics.DataLakeAnalyticsManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
