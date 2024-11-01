
import com.azure.resourcemanager.streamanalytics.models.BlobOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.BlobWriteMode;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import java.util.Arrays;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_Blob_CSV.json
     */
    /**
     * Sample code: Create a blob output with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createABlobOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output1623").withExistingStreamingjob("sjrg5023", "sj900")
            .withDatasource(
                new BlobOutputDataSource().withBlobPathPrefix("my/path").withBlobWriteMode(BlobWriteMode.ONCE)
                    .withStorageAccounts(Arrays.asList(
                        new StorageAccount().withAccountName("someAccountName").withAccountKey("fakeTokenPlaceholder")))
                    .withContainer("state").withPathPattern("{date}/{time}").withDateFormat("yyyy/MM/dd")
                    .withTimeFormat("HH"))
            .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8)).create();
    }
}
