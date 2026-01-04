
import com.azure.resourcemanager.cosmos.models.RegionForOnlineOffline;

/**
 * Samples for DatabaseAccounts OnlineRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDatabaseAccountOnlineRegion.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountOnlineRegion.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountOnlineRegion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getDatabaseAccounts().onlineRegion("rg1", "ddb1",
            new RegionForOnlineOffline().withRegion("North Europe"), com.azure.core.util.Context.NONE);
    }
}
