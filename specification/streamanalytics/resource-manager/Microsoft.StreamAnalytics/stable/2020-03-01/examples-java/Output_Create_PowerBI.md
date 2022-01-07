Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.PowerBIOutputDataSource;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_PowerBI.json
     */
    /**
     * Sample code: Create a Power BI output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAPowerBIOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output3022")
            .withExistingStreamingjob("sjrg7983", "sj2331")
            .withDatasource(
                new PowerBIOutputDataSource()
                    .withDataset("someDataset")
                    .withTable("someTable")
                    .withGroupId("ac40305e-3e8d-43ac-8161-c33799f43e95")
                    .withGroupName("MyPowerBIGroup")
                    .withRefreshToken("someRefreshToken==")
                    .withTokenUserPrincipalName("bobsmith@contoso.com")
                    .withTokenUserDisplayName("Bob Smith"))
            .create();
    }
}
```
