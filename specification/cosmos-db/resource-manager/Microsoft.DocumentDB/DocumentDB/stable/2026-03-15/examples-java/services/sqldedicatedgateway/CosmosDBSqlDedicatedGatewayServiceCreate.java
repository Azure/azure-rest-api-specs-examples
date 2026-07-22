
import com.azure.resourcemanager.cosmos.models.DedicatedGatewayType;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;
import com.azure.resourcemanager.cosmos.models.SqlDedicatedGatewayServiceResourceCreateUpdateParameters;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/services/sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceCreate.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceCreate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void sqlDedicatedGatewayServiceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().create("rg1", "ddb1", "SqlDedicatedGateway",
            new ServiceResourceCreateUpdateParameters().withProperties(
                new SqlDedicatedGatewayServiceResourceCreateUpdateParameters().withInstanceSize(ServiceSize.COSMOS_D4S)
                    .withInstanceCount(1).withDedicatedGatewayType(DedicatedGatewayType.INTEGRATED_CACHE)),
            com.azure.core.util.Context.NONE);
    }
}
