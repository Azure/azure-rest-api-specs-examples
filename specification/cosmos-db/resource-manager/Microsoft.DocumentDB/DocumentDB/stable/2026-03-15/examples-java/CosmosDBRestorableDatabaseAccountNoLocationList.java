
/**
 * Samples for RestorableDatabaseAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableDatabaseAccountNoLocationList.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountNoLocationList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBRestorableDatabaseAccountNoLocationList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableDatabaseAccounts().list(com.azure.core.util.Context.NONE);
    }
}
