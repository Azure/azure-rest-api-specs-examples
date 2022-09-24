import com.azure.core.util.Context;

/** Samples for RestorableDatabaseAccounts GetByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBRestorableDatabaseAccountGet.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableDatabaseAccountGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableDatabaseAccounts()
            .getByLocationWithResponse("West US", "d9b26648-2f53-4541-b3d8-3044f4f9810d", Context.NONE);
    }
}
