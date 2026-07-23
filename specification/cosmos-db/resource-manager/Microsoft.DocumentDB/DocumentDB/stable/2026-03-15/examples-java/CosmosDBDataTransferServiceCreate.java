
import com.azure.resourcemanager.cosmos.models.DataTransferServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDataTransferServiceCreate.json
     */
    /**
     * Sample code: DataTransferServiceCreate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void dataTransferServiceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().create("rg1", "ddb1", "DataTransfer",
            new ServiceResourceCreateUpdateParameters()
                .withProperties(new DataTransferServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S).withInstanceCount(1)),
            com.azure.core.util.Context.NONE);
    }
}
