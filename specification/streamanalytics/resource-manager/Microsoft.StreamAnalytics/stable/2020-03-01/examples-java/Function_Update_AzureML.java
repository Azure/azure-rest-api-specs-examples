import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningWebServiceFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.Function;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;

/** Samples for Functions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_AzureML.json
     */
    /**
     * Sample code: Update an Azure ML function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAnAzureMLFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Function resource =
            manager.functions().getWithResponse("sjrg7", "sj9093", "function588", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new ScalarFunctionProperties()
                    .withBinding(new AzureMachineLearningWebServiceFunctionBinding().withBatchSize(5000)))
            .apply();
    }
}
