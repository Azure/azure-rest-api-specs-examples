
import com.azure.resourcemanager.streamanalytics.models.BlobStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;
import java.util.Arrays;

/**
 * Samples for Inputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Create_Stream_Blob_CSV.json
     */
    /**
     * Sample code: Create a stream blob input with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAStreamBlobInputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().define("input8899").withExistingStreamingjob("sjrg8161", "sj6695")
            .withProperties(new StreamInputProperties()
                .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8))
                .withDatasource(new BlobStreamInputDataSource().withSourcePartitionCount(16)
                    .withStorageAccounts(Arrays.asList(
                        new StorageAccount().withAccountName("someAccountName").withAccountKey("fakeTokenPlaceholder")))
                    .withContainer("state").withPathPattern("{date}/{time}").withDateFormat("yyyy/MM/dd")
                    .withTimeFormat("HH")))
            .create();
    }
}
