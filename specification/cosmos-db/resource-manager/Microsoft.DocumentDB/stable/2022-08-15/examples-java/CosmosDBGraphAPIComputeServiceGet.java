import com.azure.core.util.Context;

/** Samples for Service Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBGraphAPIComputeServiceGet.json
     */
    /**
     * Sample code: GraphAPIComputeServiceGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void graphAPIComputeServiceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getServices()
            .getWithResponse("rg1", "ddb1", "GraphAPICompute", Context.NONE);
    }
}
