
import com.azure.resourcemanager.cosmos.models.ContinuousBackupRestoreLocation;

/**
 * Samples for TableResources RetrieveContinuousBackupInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBTableBackupInformation.json
     */
    /**
     * Sample code: CosmosDBTableCollectionBackupInformation.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBTableCollectionBackupInformation(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getTableResources().retrieveContinuousBackupInformation("rgName", "ddb1", "tableName1",
            new ContinuousBackupRestoreLocation().withLocation("North Europe"), com.azure.core.util.Context.NONE);
    }
}
