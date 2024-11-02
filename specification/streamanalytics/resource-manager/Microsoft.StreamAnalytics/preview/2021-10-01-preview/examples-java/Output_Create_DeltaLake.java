
import com.azure.resourcemanager.streamanalytics.models.BlobOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.DeltaSerialization;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import java.util.Arrays;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_DeltaLake.json
     */
    /**
     * Sample code: Create a Delta Lake output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createADeltaLakeOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output1221").withExistingStreamingjob("sjrg", "sjName")
            .withDatasource(new BlobOutputDataSource()
                .withStorageAccounts(Arrays.asList(
                    new StorageAccount().withAccountName("someAccountName").withAccountKey("fakeTokenPlaceholder")))
                .withContainer("deltaoutput"))
            .withSerialization(new DeltaSerialization().withDeltaTablePath("/folder1/table1")
                .withPartitionColumns(Arrays.asList("column1")))
            .create();
    }
}
