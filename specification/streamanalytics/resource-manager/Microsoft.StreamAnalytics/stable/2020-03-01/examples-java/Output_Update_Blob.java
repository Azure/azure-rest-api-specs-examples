import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.BlobOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Output;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_Blob.json
     */
    /**
     * Sample code: Update a blob output with CSV serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateABlobOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs().getWithResponse("sjrg5023", "sj900", "output1623", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new BlobOutputDataSource().withContainer("differentContainer"))
            .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
            .apply();
    }
}
