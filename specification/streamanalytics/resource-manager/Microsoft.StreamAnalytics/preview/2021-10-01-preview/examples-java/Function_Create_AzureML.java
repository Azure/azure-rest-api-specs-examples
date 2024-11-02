
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningStudioFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningStudioInputColumn;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningStudioInputs;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningStudioOutputColumn;
import com.azure.resourcemanager.streamanalytics.models.FunctionInput;
import com.azure.resourcemanager.streamanalytics.models.FunctionOutput;
import com.azure.resourcemanager.streamanalytics.models.ScalarFunctionProperties;
import java.util.Arrays;

/**
 * Samples for Functions CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Function_Create_AzureML.json
     */
    /**
     * Sample code: Create an Azure ML function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnAzureMLFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().define("function588").withExistingStreamingjob("sjrg7", "sj9093")
            .withProperties(new ScalarFunctionProperties()
                .withInputs(Arrays.asList(new FunctionInput().withDataType("nvarchar(max)")))
                .withOutput(new FunctionOutput().withDataType("nvarchar(max)"))
                .withBinding(new AzureMachineLearningStudioFunctionBinding().withEndpoint("someAzureMLEndpointURL")
                    .withApiKey("fakeTokenPlaceholder")
                    .withInputs(new AzureMachineLearningStudioInputs().withName("input1")
                        .withColumnNames(Arrays.asList(new AzureMachineLearningStudioInputColumn().withName("tweet")
                            .withDataType("string").withMapTo(0))))
                    .withOutputs(Arrays.asList(
                        new AzureMachineLearningStudioOutputColumn().withName("Sentiment").withDataType("string")))
                    .withBatchSize(1000)))
            .create();
    }
}
