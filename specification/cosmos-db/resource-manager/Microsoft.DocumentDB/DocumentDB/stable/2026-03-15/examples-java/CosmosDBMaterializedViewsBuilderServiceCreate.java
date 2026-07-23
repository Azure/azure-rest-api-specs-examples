
import com.azure.resourcemanager.cosmos.models.MaterializedViewsBuilderServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceResourceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ServiceSize;

/**
 * Samples for Service Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMaterializedViewsBuilderServiceCreate.json
     */
    /**
     * Sample code: MaterializedViewsBuilderServiceCreate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void materializedViewsBuilderServiceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getServices().create("rg1", "ddb1", "MaterializedViewsBuilder",
            new ServiceResourceCreateUpdateParameters()
                .withProperties(new MaterializedViewsBuilderServiceResourceCreateUpdateParameters()
                    .withInstanceSize(ServiceSize.COSMOS_D4S).withInstanceCount(1)),
            com.azure.core.util.Context.NONE);
    }
}
