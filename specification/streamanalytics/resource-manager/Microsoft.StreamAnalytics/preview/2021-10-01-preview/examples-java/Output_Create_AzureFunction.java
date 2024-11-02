
import com.azure.resourcemanager.streamanalytics.models.AzureFunctionOutputDataSource;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_AzureFunction.json
     */
    /**
     * Sample code: Create an Azure Function output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnAzureFunctionOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("azureFunction1").withExistingStreamingjob("sjrg", "sjName")
            .withDatasource(new AzureFunctionOutputDataSource().withFunctionAppName("functionappforasaautomation")
                .withFunctionName("HttpTrigger2").withMaxBatchSize(256.0F).withMaxBatchCount(100.0F))
            .create();
    }
}
