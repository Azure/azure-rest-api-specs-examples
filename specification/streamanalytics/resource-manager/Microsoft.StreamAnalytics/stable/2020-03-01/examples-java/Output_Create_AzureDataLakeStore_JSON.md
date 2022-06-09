```java
import com.azure.resourcemanager.streamanalytics.models.AzureDataLakeStoreOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.JsonOutputSerializationFormat;
import com.azure.resourcemanager.streamanalytics.models.JsonSerialization;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureDataLakeStore_JSON.json
     */
    /**
     * Sample code: Create an Azure Data Lake Store output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAnAzureDataLakeStoreOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output5195")
            .withExistingStreamingjob("sjrg6912", "sj3310")
            .withDatasource(
                new AzureDataLakeStoreOutputDataSource()
                    .withAccountName("someaccount")
                    .withTenantId("cea4e98b-c798-49e7-8c40-4a2b3beb47dd")
                    .withFilePathPrefix("{date}/{time}")
                    .withDateFormat("yyyy/MM/dd")
                    .withTimeFormat("HH")
                    .withRefreshToken("someRefreshToken==")
                    .withTokenUserPrincipalName("bobsmith@contoso.com")
                    .withTokenUserDisplayName("Bob Smith"))
            .withSerialization(
                new JsonSerialization().withEncoding(Encoding.UTF8).withFormat(JsonOutputSerializationFormat.ARRAY))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
