import com.azure.core.util.Context;

/** Samples for TableResources DeleteTable. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBTableDelete.json
     */
    /**
     * Sample code: CosmosDBTableDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBTableDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getTableResources()
            .deleteTable("rg1", "ddb1", "tableName", Context.NONE);
    }
}
