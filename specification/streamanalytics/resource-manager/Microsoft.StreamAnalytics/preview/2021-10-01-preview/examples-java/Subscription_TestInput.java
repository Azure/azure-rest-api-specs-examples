
import com.azure.resourcemanager.streamanalytics.fluent.models.InputInner;
import com.azure.resourcemanager.streamanalytics.fluent.models.TestInputInner;
import com.azure.resourcemanager.streamanalytics.models.BlobStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;
import java.util.Arrays;

/**
 * Samples for Subscriptions TestInput.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Subscription_TestInput.json
     */
    /**
     * Sample code: Test the Stream Analytics input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        testTheStreamAnalyticsInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.subscriptions().testInput("West US",
            new TestInputInner().withInput(new InputInner().withProperties(new StreamInputProperties()
                .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8))
                .withDatasource(new BlobStreamInputDataSource().withSourcePartitionCount(16)
                    .withStorageAccounts(Arrays.asList(
                        new StorageAccount().withAccountName("someAccountName").withAccountKey("fakeTokenPlaceholder")))
                    .withContainer("state").withPathPattern("{date}/{time}").withDateFormat("yyyy/MM/dd")
                    .withTimeFormat("HH")))),
            com.azure.core.util.Context.NONE);
    }
}
