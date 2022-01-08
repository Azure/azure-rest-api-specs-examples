Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.BlobReferenceInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.Input;
import com.azure.resourcemanager.streamanalytics.models.ReferenceInputProperties;

/** Samples for Inputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Reference_Blob.json
     */
    /**
     * Sample code: Update a reference blob input.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAReferenceBlobInput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Input resource = manager.inputs().getWithResponse("sjrg8440", "sj9597", "input7225", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                new ReferenceInputProperties()
                    .withSerialization(new CsvSerialization().withFieldDelimiter("|").withEncoding(Encoding.UTF8))
                    .withDatasource(new BlobReferenceInputDataSource()))
            .apply();
    }
}
```
