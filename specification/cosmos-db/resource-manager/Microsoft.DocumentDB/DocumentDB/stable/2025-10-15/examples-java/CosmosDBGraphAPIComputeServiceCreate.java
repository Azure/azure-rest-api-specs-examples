
import com.azure.resourcemanager.cosmos.models.GraphApiComputeServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGraphAPIComputeServiceCreate.json
     */
    /**
     * Sample code: GraphAPIComputeServiceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void graphAPIComputeServiceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().create("rg1", "ddb1", "GraphAPICompute",
            new ServiceResourceCreateUpdateParameters()
                .withProperties(new GraphApiComputeServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S).withInstanceCount(1)),
            com.azure.core.util.Context.NONE);
    }
}
