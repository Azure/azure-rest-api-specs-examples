
import com.azure.resourcemanager.streamanalytics.models.BlobOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.DeltaSerialization;
import com.azure.resourcemanager.streamanalytics.models.Output;
import java.util.Arrays;

/**
 * Samples for Outputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Update_DeltaLake.json
     */
    /**
     * Sample code: Update a Delta Lake output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateADeltaLakeOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs()
            .getWithResponse("sjrg", "sjName", "output1221", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDatasource(new BlobOutputDataSource().withContainer("deltaoutput2"))
            .withSerialization(new DeltaSerialization().withDeltaTablePath("/folder1/table2")
                .withPartitionColumns(Arrays.asList("column2")))
            .apply();
    }
}
