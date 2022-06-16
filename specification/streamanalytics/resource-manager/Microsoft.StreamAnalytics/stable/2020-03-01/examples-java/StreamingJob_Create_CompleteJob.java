import com.azure.resourcemanager.streamanalytics.fluent.models.InputInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.OutputInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.TransformationInner;
import com.azure.resourcemanager.streamanalytics.models.AzureSqlDatabaseOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.BlobStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CompatibilityLevel;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.EventsOutOfOrderPolicy;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;
import com.azure.resourcemanager.streamanalytics.models.OutputErrorPolicy;
import com.azure.resourcemanager.streamanalytics.models.Sku;
import com.azure.resourcemanager.streamanalytics.models.SkuName;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for StreamingJobs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
     */
    /**
     * Sample code: Create a complete streaming job (a streaming job with a transformation, at least 1 input and at
     * least 1 output).
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createACompleteStreamingJobAStreamingJobWithATransformationAtLeast1InputAndAtLeast1Output(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .streamingJobs()
            .define("sj7804")
            .withRegion("West US")
            .withExistingResourceGroup("sjrg3276")
            .withTags(mapOf("key1", "value1", "key3", "value3", "randomKey", "randomValue"))
            .withSku(new Sku().withName(SkuName.STANDARD))
            .withEventsOutOfOrderPolicy(EventsOutOfOrderPolicy.DROP)
            .withOutputErrorPolicy(OutputErrorPolicy.DROP)
            .withEventsOutOfOrderMaxDelayInSeconds(0)
            .withEventsLateArrivalMaxDelayInSeconds(5)
            .withDataLocale("en-US")
            .withCompatibilityLevel(CompatibilityLevel.ONE_ZERO)
            .withInputs(
                Arrays
                    .asList(
                        new InputInner()
                            .withProperties(
                                new StreamInputProperties()
                                    .withSerialization(new JsonSerialization().withEncoding(Encoding.UTF8))
                                    .withDatasource(
                                        new BlobStreamInputDataSource()
                                            .withStorageAccounts(
                                                Arrays
                                                    .asList(
                                                        new StorageAccount()
                                                            .withAccountName("yourAccountName")
                                                            .withAccountKey("yourAccountKey==")))
                                            .withContainer("containerName")
                                            .withPathPattern("")))
                            .withName("inputtest")))
            .withTransformation(
                new TransformationInner()
                    .withName("transformationtest")
                    .withStreamingUnits(1)
                    .withQuery("Select Id, Name from inputtest"))
            .withOutputs(
                Arrays
                    .asList(
                        new OutputInner()
                            .withName("outputtest")
                            .withDatasource(new AzureSqlDatabaseOutputDataSource())))
            .withFunctions(Arrays.asList())
            .create();
    }

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
