
/**
 * Samples for Queries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/QueryPackQueriesGet.json
     */
    /**
     * Sample code: QueryGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void queryGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.queries().getWithResponse("my-resource-group", "my-querypack", "a449f8af-8e64-4b3a-9b16-5a7165ff98c4",
            com.azure.core.util.Context.NONE);
    }
}
