
import com.azure.resourcemanager.mongocluster.models.DataApiMode;
import com.azure.resourcemanager.mongocluster.models.DataApiProperties;
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_PatchDataApi.json
     */
    /**
     * Sample code: Enables data API on a mongo cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        enablesDataAPIOnAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(
                new MongoClusterUpdateProperties().withDataApi(new DataApiProperties().withMode(DataApiMode.ENABLED)))
            .apply();
    }
}
