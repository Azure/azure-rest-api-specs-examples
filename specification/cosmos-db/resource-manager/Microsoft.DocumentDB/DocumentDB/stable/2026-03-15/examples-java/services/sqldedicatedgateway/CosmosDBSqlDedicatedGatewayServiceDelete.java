
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceDelete.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void sqlDedicatedGatewayServiceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().delete("rg1", "ddb1", "SqlDedicatedGateway",
            com.azure.core.util.Context.NONE);
    }
}
