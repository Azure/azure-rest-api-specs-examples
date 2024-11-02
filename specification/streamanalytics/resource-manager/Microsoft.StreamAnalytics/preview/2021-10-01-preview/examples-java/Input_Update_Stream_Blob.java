
import com.azure.resourcemanager.streamanalytics.models.BlobStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Input;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;

/**
 * Samples for Inputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Update_Stream_Blob.json
     */
    /**
     * Sample code: Update a stream blob input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAStreamBlobInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Input resource = manager.inputs()
            .getWithResponse("sjrg8161", "sj6695", "input8899", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new StreamInputProperties()
                .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
                .withDatasource(new BlobStreamInputDataSource().withSourcePartitionCount(32)))
            .apply();
    }
}
