
import com.azure.resourcemanager.cosmos.models.DataTransferServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDataTransferServiceCreate.json
     */
    /**
     * Sample code: DataTransferServiceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dataTransferServiceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().create("rg1", "ddb1", "DataTransfer",
            new ServiceResourceCreateUpdateParameters()
                .withProperties(new DataTransferServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S).withInstanceCount(1)),
            com.azure.core.util.Context.NONE);
    }
}
