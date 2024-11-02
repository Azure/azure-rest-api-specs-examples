
import com.azure.resourcemanager.streamanalytics.models.BlobReferenceInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.ReferenceInputProperties;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import java.util.Arrays;

/**
 * Samples for Inputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Create_Reference_Blob_CSV.json
     */
    /**
     * Sample code: Create a reference blob input with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAReferenceBlobInputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().define("input7225").withExistingStreamingjob("sjrg8440", "sj9597")
            .withProperties(new ReferenceInputProperties()
                .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8))
                .withDatasource(new BlobReferenceInputDataSource().withBlobName("myblobinput")
                    .withDeltaPathPattern("/testBlob/{date}/delta/{time}/").withSourcePartitionCount(16)
                    .withFullSnapshotRefreshRate("16:14:30").withDeltaSnapshotRefreshRate("16:14:30")
                    .withStorageAccounts(Arrays.asList(
                        new StorageAccount().withAccountName("someAccountName").withAccountKey("fakeTokenPlaceholder")))
                    .withContainer("state").withPathPattern("{date}/{time}").withDateFormat("yyyy/MM/dd")
                    .withTimeFormat("HH")))
            .create();
    }
}
