
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningServiceFunctionBinding;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningServiceInputColumn;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningServiceOutputColumn;
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
     * Function_Create_AzureMLService.json
     */
    /**
     * Sample code: Create an Azure ML Service function.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnAzureMLServiceFunction(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.functions().define("function588").withExistingStreamingjob("sjrg", "sjName")
            .withProperties(new ScalarFunctionProperties()
                .withInputs(Arrays.asList(new FunctionInput().withDataType("nvarchar(max)")))
                .withOutput(new FunctionOutput().withDataType("nvarchar(max)"))
                .withBinding(new AzureMachineLearningServiceFunctionBinding().withEndpoint("someAzureMLEndpointURL")
                    .withApiKey("fakeTokenPlaceholder")
                    .withInputs(Arrays.asList(new AzureMachineLearningServiceInputColumn().withName("data")
                        .withDataType("array").withMapTo(0)))
                    .withOutputs(Arrays.asList(
                        new AzureMachineLearningServiceOutputColumn().withName("Sentiment").withDataType("string")))
                    .withBatchSize(1000).withNumberOfParallelRequests(1).withInputRequestName("Inputs")
                    .withOutputResponseName("Results")))
            .create();
    }
}
