import com.azure.core.util.Context;

/** Samples for DatabaseAccounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseAccountListByResourceGroup.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountListByResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .listByResourceGroup("rg1", Context.NONE);
    }
}
