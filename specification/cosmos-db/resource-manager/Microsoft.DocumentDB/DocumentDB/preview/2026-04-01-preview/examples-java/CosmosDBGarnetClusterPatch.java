
import com.azure.resourcemanager.cosmos.models.ClusterType;
import com.azure.resourcemanager.cosmos.models.GarnetAuthenticationType;
import com.azure.resourcemanager.cosmos.models.GarnetClusterResourcePatch;
import com.azure.resourcemanager.cosmos.models.GarnetClusterResourcePatchProperties;

/**
 * Samples for GarnetClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBGarnetClusterPatch.json
     */
    /**
     * Sample code: CosmosDBGarnetClusterPatch.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBGarnetClusterPatch(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getGarnetClusters().update("garnet-prod-rg", "garnet-prod",
            new GarnetClusterResourcePatch()
                .withProperties(new GarnetClusterResourcePatchProperties().withClusterType(ClusterType.PRODUCTION)
                    .withAuthenticationMethod(GarnetAuthenticationType.ENTRA).withPersistence(true)),
            com.azure.core.util.Context.NONE);
    }
}
