import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;
import com.azure.resourcemanager.cosmos.models.ServiceType;

/** Samples for Service Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBGraphAPIComputeServiceCreate.json
     */
    /**
     * Sample code: GraphAPIComputeServiceCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void graphAPIComputeServiceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .create(
                "rg1",
                "ddb1",
                "GraphAPICompute",
                new ServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S)
                    .withInstanceCount(1)
                    .withServiceType(ServiceType.GRAPH_APICOMPUTE),
                com.azure.core.util.Context.NONE);
    }
}
