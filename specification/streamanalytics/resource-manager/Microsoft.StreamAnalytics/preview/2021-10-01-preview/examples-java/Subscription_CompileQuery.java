
import com.azure.resourcemanager.streamanalytics.models.CompatibilityLevel;
import com.azure.resourcemanager.streamanalytics.models.CompileQuery;
import com.azure.resourcemanager.streamanalytics.models.FunctionInput;
import com.azure.resourcemanager.streamanalytics.models.FunctionOutput;
import com.azure.resourcemanager.streamanalytics.models.JobType;
import com.azure.resourcemanager.streamanalytics.models.QueryFunction;
import com.azure.resourcemanager.streamanalytics.models.QueryInput;
import java.util.Arrays;

/**
 * Samples for Subscriptions CompileQuery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Subscription_CompileQuery.json
     */
    /**
     * Sample code: Compile the Stream Analytics query.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        compileTheStreamAnalyticsQuery(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.subscriptions().compileQueryWithResponse("West US",
            new CompileQuery().withQuery("SELECT\r\n    *\r\nINTO\r\n    [output1]\r\nFROM\r\n    [input1]")
                .withInputs(Arrays.asList(new QueryInput().withName("input1").withType("Stream")))
                .withFunctions(Arrays.asList(new QueryFunction().withName("function1").withType("Scalar")
                    .withBindingType("Microsoft.StreamAnalytics/JavascriptUdf")
                    .withInputs(Arrays.asList(new FunctionInput().withDataType("any")))
                    .withOutput(new FunctionOutput().withDataType("bigint"))))
                .withJobType(JobType.CLOUD).withCompatibilityLevel(CompatibilityLevel.ONE_TWO),
            com.azure.core.util.Context.NONE);
    }
}
