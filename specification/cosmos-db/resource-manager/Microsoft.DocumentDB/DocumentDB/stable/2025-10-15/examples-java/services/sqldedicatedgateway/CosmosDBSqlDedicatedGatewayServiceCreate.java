
import com.azure.resourcemanager.cosmos.models.DedicatedGatewayType;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;
import com.azure.resourcemanager.cosmos.models.SqlDedicatedGatewayServiceResourceCreateUpdateParameters;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/services/
     * sqldedicatedgateway/CosmosDBSqlDedicatedGatewayServiceCreate.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sqlDedicatedGatewayServiceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().create("rg1", "ddb1", "SqlDedicatedGateway",
            new ServiceResourceCreateUpdateParameters().withProperties(
                new SqlDedicatedGatewayServiceResourceCreateUpdateParameters().withInstanceSize(ServiceSize.COSMOS_D4S)
                    .withInstanceCount(1).withDedicatedGatewayType(DedicatedGatewayType.INTEGRATED_CACHE)),
            com.azure.core.util.Context.NONE);
    }
}
