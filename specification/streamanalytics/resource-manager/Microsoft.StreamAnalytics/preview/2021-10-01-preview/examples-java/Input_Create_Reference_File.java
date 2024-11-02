
import com.azure.resourcemanager.streamanalytics.models.FileReferenceInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.ReferenceInputProperties;

/**
 * Samples for Inputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Create_Reference_File.json
     */
    /**
     * Sample code: Create a reference file input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAReferenceFileInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().define("input7225").withExistingStreamingjob("sjrg8440", "sj9597")
            .withProperties(
                new ReferenceInputProperties().withDatasource(new FileReferenceInputDataSource().withPath("my/path")))
            .create();
    }
}
