
import com.azure.resourcemanager.streamanalytics.models.DocumentDbOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Output;

/**
 * Samples for Outputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Update_DocumentDB.json
     */
    /**
     * Sample code: Update a DocumentDB output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateADocumentDBOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs()
            .getWithResponse("sjrg7983", "sj2331", "output3022", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDatasource(new DocumentDbOutputDataSource().withPartitionKey("fakeTokenPlaceholder"))
            .apply();
    }
}
