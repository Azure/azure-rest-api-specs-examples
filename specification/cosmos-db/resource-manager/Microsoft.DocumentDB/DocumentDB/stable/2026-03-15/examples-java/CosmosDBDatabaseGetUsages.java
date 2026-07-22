
/**
 * Samples for Database ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseGetUsages.json
     */
    /**
     * Sample code: CosmosDBDatabaseGetUsages.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseGetUsages(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabases().listUsages("rg1", "ddb1", "databaseRid", "name.value eq 'Storage'",
            com.azure.core.util.Context.NONE);
    }
}
