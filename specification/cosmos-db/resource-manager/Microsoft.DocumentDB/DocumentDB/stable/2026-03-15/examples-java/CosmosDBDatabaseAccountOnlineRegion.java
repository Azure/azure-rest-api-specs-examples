
import com.azure.resourcemanager.cosmos.models.RegionForOnlineOffline;

/**
 * Samples for DatabaseAccounts OnlineRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountOnlineRegion.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountOnlineRegion.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountOnlineRegion(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().onlineRegion("rg1", "ddb1",
            new RegionForOnlineOffline().withRegion("North Europe"), com.azure.core.util.Context.NONE);
    }
}
