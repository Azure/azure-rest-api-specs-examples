
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for TableResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBTableBackupInformation.json
     */
    /**
     * Sample code: CosmosDBTableCollectionBackupInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableCollectionBackupInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getTableResources().retrieveContinuousBackupInformation(
            "rgName", "ddb1", "tableName1", new ContinuousBackupRestoreLocation().withLocation("North Europe"),
            com.azure.core.util.Context.NONE);
    }
}
