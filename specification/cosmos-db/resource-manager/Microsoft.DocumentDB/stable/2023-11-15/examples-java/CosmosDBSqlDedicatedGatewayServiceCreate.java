
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;
import com.azure.resourcemanager.cosmos.models.ServiceType;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/
     * CosmosDBSqlDedicatedGatewayServiceCreate.json
     */
    /**
     * Sample code: SqlDedicatedGatewayServiceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void sqlDedicatedGatewayServiceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices()
            .create("rg1", "ddb1", "SqlDedicatedGateway",
                new ServiceResourceCreateUpdateParameters().withInstanceSize(ServiceSize.COSMOS_D4S)
                    .withInstanceCount(1).withServiceType(ServiceType.SQL_DEDICATED_GATEWAY),
                com.azure.core.util.Context.NONE);
    }
}
