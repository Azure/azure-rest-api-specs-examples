import com.azure.resourcemanager.digitaltwins.models.AzureDataExplorerConnectionProperties;
import com.azure.resourcemanager.digitaltwins.models.IdentityType;
import com.azure.resourcemanager.digitaltwins.models.ManagedIdentityReference;

/** Samples for TimeSeriesDatabaseConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/TimeSeriesDatabaseConnectionsPut_WithUserIdentity_example.json
     */
    /**
     * Sample code: Create or replace a time series database connection for a DigitalTwins instance with user assigned
     * identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void createOrReplaceATimeSeriesDatabaseConnectionForADigitalTwinsInstanceWithUserAssignedIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .timeSeriesDatabaseConnections()
            .define("myConnection")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(
                new AzureDataExplorerConnectionProperties()
                    .withIdentity(
                        new ManagedIdentityReference()
                            .withType(IdentityType.USER_ASSIGNED)
                            .withUserAssignedIdentity(
                                "/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity"))
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
