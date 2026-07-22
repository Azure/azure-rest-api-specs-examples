
/**
 * Samples for RestorableDatabaseAccounts GetByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableDatabaseAccountGet.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableDatabaseAccountGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableDatabaseAccounts().getByLocationWithResponse("West US",
            "d9b26648-2f53-4541-b3d8-3044f4f9810d", com.azure.core.util.Context.NONE);
    }
}
