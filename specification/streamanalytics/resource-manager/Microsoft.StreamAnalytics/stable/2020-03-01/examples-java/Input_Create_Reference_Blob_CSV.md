Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.BlobReferenceInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.ReferenceInputProperties;

/** Samples for Inputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Reference_Blob_CSV.json
     */
    /**
     * Sample code: Create a reference blob input with CSV serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAReferenceBlobInputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .inputs()
            .define("input7225")
            .withExistingStreamingjob("sjrg8440", "sj9597")
            .withProperties(
                new ReferenceInputProperties()
                    .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8))
                    .withDatasource(new BlobReferenceInputDataSource()))
            .create();
    }
}
```
