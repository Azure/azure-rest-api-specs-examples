
/**
 * Samples for Transformations CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Transformation_Create.json
     */
    /**
     * Sample code: Create a transformation.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createATransformation(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.transformations().define("transformation952").withExistingStreamingjob("sjrg4423", "sj8374")
            .withStreamingUnits(6).withQuery("Select Id, Name from inputtest").create();
    }
}
