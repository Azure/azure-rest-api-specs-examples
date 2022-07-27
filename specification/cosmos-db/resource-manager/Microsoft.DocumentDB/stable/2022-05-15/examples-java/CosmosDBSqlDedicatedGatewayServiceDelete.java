import com.azure.core.util.Context;

/** Samples for Service Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlDedicatedGatewayServiceDelete.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sqlDedicatedGatewayServiceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .delete("rg1", "ddb1", "SqlDedicatedGateway", Context.NONE);
    }
}
