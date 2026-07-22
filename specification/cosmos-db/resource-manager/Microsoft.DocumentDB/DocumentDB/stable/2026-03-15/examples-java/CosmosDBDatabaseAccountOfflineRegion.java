
import com.azure.resourcemanager.cosmos.models.RegionForOnlineOffline;

/**
 * Samples for DatabaseAccounts OfflineRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountOfflineRegion.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountOfflineRegion.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountOfflineRegion(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().offlineRegion("rg1", "ddb1",
            new RegionForOnlineOffline().withRegion("North Europe"), com.azure.core.util.Context.NONE);
    }
}
