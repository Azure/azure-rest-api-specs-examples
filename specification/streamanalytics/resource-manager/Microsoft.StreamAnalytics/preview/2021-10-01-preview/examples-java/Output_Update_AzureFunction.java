
import com.azure.resourcemanager.streamanalytics.models.AzureFunctionOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Output;

/**
 * Samples for Outputs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Update_AzureFunction.json
     */
    /**
     * Sample code: Update an Azure Function output.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAnAzureFunctionOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs()
            .getWithResponse("sjrg", "sjName", "azureFunction1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDatasource(new AzureFunctionOutputDataSource().withFunctionName("differentFunctionName"))
            .apply();
    }
}
