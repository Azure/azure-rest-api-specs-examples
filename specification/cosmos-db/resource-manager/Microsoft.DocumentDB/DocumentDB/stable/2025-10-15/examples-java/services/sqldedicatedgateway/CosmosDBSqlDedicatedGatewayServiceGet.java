
/**
 * Samples for Service Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/services/
     * sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceGet.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sqlDedicatedGatewayServiceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().getWithResponse("rg1", "ddb1",
            "SqlDedicatedGateway", com.azure.core.util.Context.NONE);
    }
}
