
/**
 * Samples for RestorableSqlResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBRestorableSqlResourceList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlResourceList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBRestorableSqlResourceList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getRestorableSqlResources().list("WestUS", "d9b26648-2f53-4541-b3d8-3044f4f9810d",
            "WestUS", "06/01/2022 4:56", com.azure.core.util.Context.NONE);
    }
}
