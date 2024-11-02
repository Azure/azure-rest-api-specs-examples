
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningStudioFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.Function;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;

/**
 * Samples for Functions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Update_AzureML.json
     */
    /**
     * Sample code: Update an Azure ML function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        updateAnAzureMLFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Function resource = manager.functions()
            .getWithResponse("sjrg7", "sj9093", "function588", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ScalarFunctionProperties()
            .withBinding(new AzureMachineLearningStudioFunctionBinding().withBatchSize(5000))).apply();
    }
}
