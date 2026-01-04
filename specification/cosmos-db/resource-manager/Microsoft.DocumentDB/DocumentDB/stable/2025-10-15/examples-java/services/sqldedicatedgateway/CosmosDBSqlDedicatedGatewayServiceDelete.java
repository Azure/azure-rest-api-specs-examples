
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/services/
     * sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceDelete.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sqlDedicatedGatewayServiceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().delete("rg1", "ddb1", "SqlDedicatedGateway",
            com.azure.core.util.Context.NONE);
    }
}
