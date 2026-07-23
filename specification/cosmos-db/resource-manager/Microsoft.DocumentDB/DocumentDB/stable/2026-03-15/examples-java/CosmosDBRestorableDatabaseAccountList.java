
/**
 * Samples for RestorableDatabaseAccounts ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableDatabaseAccountList.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableDatabaseAccountList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableDatabaseAccounts().listByLocation("West US",
            com.azure.core.util.Context.NONE);
    }
}
