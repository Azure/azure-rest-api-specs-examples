```java
import com.azure.core.util.Context;

/** Samples for StreamingJobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Get_NoExpand.json
     */
    /**
     * Sample code: Get a streaming job and do not use the $expand OData query parameter.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamingJobAndDoNotUseTheExpandODataQueryParameter(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.streamingJobs().getByResourceGroupWithResponse("sjrg6936", "sj59", null, Context.NONE);
    }

    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Get_Expand.json
     */
    /**
     * Sample code: Get a streaming job and use the $expand OData query parameter to expand inputs, outputs,
     * transformation, and functions.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        getAStreamingJobAndUseTheExpandODataQueryParameterToExpandInputsOutputsTransformationAndFunctions(
            com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .streamingJobs()
            .getByResourceGroupWithResponse(
                "sjrg3276", "sj7804", "inputs,outputs,transformation,functions", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
