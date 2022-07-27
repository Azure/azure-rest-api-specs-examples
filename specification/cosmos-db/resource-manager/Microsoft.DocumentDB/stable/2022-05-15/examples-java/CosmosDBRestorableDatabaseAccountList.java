import com.azure.core.util.Context;

/** Samples for RestorableDatabaseAccounts ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBRestorableDatabaseAccountList.json
     */
    /**
     * Sample code: CosmosDBRestorableDatabaseAccountList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableDatabaseAccountList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableDatabaseAccounts()
            .listByLocation("West US", Context.NONE);
    }
}
