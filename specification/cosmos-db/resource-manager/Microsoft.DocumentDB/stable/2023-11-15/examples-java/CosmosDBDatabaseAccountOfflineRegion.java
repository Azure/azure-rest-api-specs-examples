
import com.azure.resourcemanager.cosmos.models.RegionForOnlineOffline;

/**
 * Samples for DatabaseAccounts OfflineRegion.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/
     * CosmosDBDatabaseAccountOfflineRegion.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountOfflineRegion.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountOfflineRegion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().offlineRegion("rg1", "ddb1",
            new RegionForOnlineOffline(), com.azure.core.util.Context.NONE);
    }
}
