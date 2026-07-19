
/**
 * Samples for SyncGroups ListHubSchemas.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncGroupGetHubSchema.json
     */
    /**
     * Sample code: Get a hub database schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAHubDatabaseSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().listHubSchemas("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
