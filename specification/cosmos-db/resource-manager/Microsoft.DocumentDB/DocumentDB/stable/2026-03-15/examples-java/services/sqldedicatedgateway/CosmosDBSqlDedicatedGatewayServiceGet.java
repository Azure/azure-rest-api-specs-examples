
/**
 * Samples for Service Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceGet.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void sqlDedicatedGatewayServiceGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().getWithResponse("rg1", "ddb1", "SqlDedicatedGateway",
            com.azure.core.util.Context.NONE);
    }
}
