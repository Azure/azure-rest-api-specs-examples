
import com.azure.resourcemanager.streamanalytics.fluent.models.InputInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.OutputInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.StreamingJobInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.TestQueryInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.TransformationInner;
import com.azure.resourcemanager.streamanalytics.models.CompatibilityLevel;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventsOutOfOrderPolicy;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.OutputErrorPolicy;
import com.azure.resourcemanager.streamanalytics.models.RawOutputDatasource;
import com.azure.resourcemanager.streamanalytics.models.RawStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Sku;
import com.azure.resourcemanager.streamanalytics.models.SkuName;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;
import com.azure.resourcemanager.streamanalytics.models.TestQueryDiagnostics;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Subscriptions TestQuery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Subscription_TestQuery.json
     */
    /**
     * Sample code: Test the Stream Analytics query.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        testTheStreamAnalyticsQuery(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.subscriptions()
            .testQuery("West US",
                new TestQueryInner()
                    .withDiagnostics(new TestQueryDiagnostics().withWriteUri(
                        "http://myoutput.com").withPath(
                            "/pathto/subdirectory"))
                    .withStreamingJob(
                        new StreamingJobInner().withLocation("West US")
                            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key3", "fakeTokenPlaceholder", "randomKey",
                                "fakeTokenPlaceholder"))
                            .withSkuPropertiesSku(new Sku().withName(SkuName.STANDARD))
                            .withEventsOutOfOrderPolicy(EventsOutOfOrderPolicy.DROP)
                            .withOutputErrorPolicy(OutputErrorPolicy.DROP).withEventsOutOfOrderMaxDelayInSeconds(0)
                            .withEventsLateArrivalMaxDelayInSeconds(5).withDataLocale("en-US")
                            .withCompatibilityLevel(CompatibilityLevel.ONE_ZERO)
                            .withInputs(Arrays.asList(new InputInner().withProperties(new StreamInputProperties()
                                .withSerialization(new JsonSerialization().withEncoding(Encoding.UTF8))
                                .withDatasource(new RawStreamInputDataSource().withPayloadUri("http://myinput.com")))
                                .withName("inputtest")))
                            .withTransformation(new TransformationInner().withName("transformationtest")
                                .withStreamingUnits(1).withQuery("Select Id, Name from inputtest"))
                            .withOutputs(Arrays.asList(new OutputInner().withName("outputtest")
                                .withDatasource(new RawOutputDatasource().withPayloadUri("http://myoutput.com"))
                                .withSerialization(new JsonSerialization())))
                            .withFunctions(Arrays.asList())),
                com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
