Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.DocumentDbOutputDataSource;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_DocumentDB.json
     */
    /**
     * Sample code: Create a DocumentDB output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createADocumentDBOutput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output3022")
            .withExistingStreamingjob("sjrg7983", "sj2331")
            .withDatasource(
                new DocumentDbOutputDataSource()
                    .withAccountId("someAccountId")
                    .withAccountKey("accountKey==")
                    .withDatabase("db01")
                    .withCollectionNamePattern("collection")
                    .withPartitionKey("key")
                    .withDocumentId("documentId"))
            .create();
    }
}
```
