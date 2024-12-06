
import com.azure.resourcemanager.digitaltwins.models.AzureDataExplorerConnectionProperties;
import com.azure.resourcemanager.digitaltwins.models.RecordPropertyAndItemRemovals;

/**
 * Samples for TimeSeriesDatabaseConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * TimeSeriesDatabaseConnectionsPut_example.json
     */
    /**
     * Sample code: Create or replace a time series database connection for a DigitalTwins instance.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstance(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.timeSeriesDatabaseConnections().define("myConnection")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(new AzureDataExplorerConnectionProperties().withAdxResourceId(
                "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster")
                .withAdxEndpointUri("https://mycluster.kusto.windows.net").withAdxDatabaseName("myDatabase")
                .withAdxTableName("myPropertyUpdatesTable")
                .withAdxTwinLifecycleEventsTableName("myTwinLifecycleEventsTable")
                .withAdxRelationshipLifecycleEventsTableName("myRelationshipLifecycleEventsTable")
                .withEventHubEndpointUri("sb://myeh.servicebus.windows.net/").withEventHubEntityPath("myeh")
                .withEventHubNamespaceResourceId(
                    "/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh")
                .withRecordPropertyAndItemRemovals(RecordPropertyAndItemRemovals.TRUE))
            .create();
    }
}
