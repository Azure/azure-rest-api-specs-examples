import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters;
import com.azure.resourcemanager.streamanalytics.models.UdfType;

/** Samples for Functions RetrieveDefaultDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_RetrieveDefaultDefinition_AzureML.json
     */
    /**
     * Sample code: Retrieve the default definition for an Azure ML function.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void retrieveTheDefaultDefinitionForAnAzureMLFunction(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .functions()
            .retrieveDefaultDefinitionWithResponse(
                "sjrg7",
                "sj9093",
                "function588",
                new AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters()
                    .withExecuteEndpoint("someAzureMLExecuteEndpointUrl")
                    .withUdfType(UdfType.SCALAR),
                Context.NONE);
    }
}
