
import com.azure.resourcemanager.cosmos.models.GraphApiComputeServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBGraphAPIComputeServiceCreate.json
     */
    /**
     * Sample code: GraphAPIComputeServiceCreate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void graphAPIComputeServiceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().create("rg1", "ddb1", "GraphAPICompute",
            new ServiceResourceCreateUpdateParameters()
                .withProperties(new GraphApiComputeServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S).withInstanceCount(1)),
            com.azure.core.util.Context.NONE);
    }
}
