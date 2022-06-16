import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.AzureDataLakeStoreOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.Output;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_AzureDataLakeStore.json
     */
    /**
     * Sample code: Update an Azure Data Lake Store output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAnAzureDataLakeStoreOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource =
            manager.outputs().getWithResponse("sjrg6912", "sj3310", "output5195", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new AzureDataLakeStoreOutputDataSource().withAccountName("differentaccount"))
            .withSerialization(
                new JsonSerialization()
                    .withEncoding(Encoding.UTF8)
                    .withFormat(JsonOutputSerializationFormat.LINE_SEPARATED))
            .apply();
    }
}
