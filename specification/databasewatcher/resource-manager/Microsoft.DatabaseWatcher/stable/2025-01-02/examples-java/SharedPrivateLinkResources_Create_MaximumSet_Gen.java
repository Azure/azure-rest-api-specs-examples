
import com.azure.resourcemanager.databasewatcher.models.SharedPrivateLinkResourceProperties;

/**
 * Samples for SharedPrivateLinkResources Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: SharedPrivateLinkResources_Create_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void sharedPrivateLinkResourcesCreateMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.sharedPrivateLinkResources().define("monitoringh22eed")
            .withExistingWatcher("apiTest-ddat4p", "databasemo3ej9ih")
            .withProperties(new SharedPrivateLinkResourceProperties().withPrivateLinkResourceId(
                "/subscriptions/49e0fbd3-75e8-44e7-96fd-5b64d9ad818d/resourceGroups/apiTest-ddat4p/providers/Microsoft.KeyVault/vaults/kvmo3ej9ih")
                .withGroupId("vault").withRequestMessage("request message").withDnsZone("ec3ae9d410ba"))
            .create();
    }
}
