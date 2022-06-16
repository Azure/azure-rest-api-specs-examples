import com.azure.core.util.Context;

/** Samples for DatabaseAccounts CheckNameExists. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseAccountCheckNameExists.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountCheckNameExists.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseAccountCheckNameExists(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabaseAccounts()
            .checkNameExistsWithResponse("ddb1", Context.NONE);
    }
}
