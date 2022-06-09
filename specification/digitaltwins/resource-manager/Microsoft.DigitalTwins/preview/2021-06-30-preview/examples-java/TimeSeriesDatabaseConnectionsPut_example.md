```java
import com.azure.resourcemanager.digitaltwins.models.AzureDataExplorerConnectionProperties;

/** Samples for TimeSeriesDatabaseConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsPut_example.json
     */
    /**
     * Sample code: Create or replace a time series database connection for a DigitalTwins instance.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .timeSeriesDatabaseConnections()
            .define("myConnection")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(
                new AzureDataExplorerConnectionProperties()
                    .withAdxResourceId(
                        "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster")
                    .withAdxEndpointUri("https://mycluster.kusto.windows.net")
                    .withAdxDatabaseName("myDatabase")
                    .withAdxTableName("myTable")
                    .withEventHubEndpointUri("sb://myeh.servicebus.windows.net/")
                    .withEventHubEntityPath("myeh")
                    .withEventHubNamespaceResourceId(
                        "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.
