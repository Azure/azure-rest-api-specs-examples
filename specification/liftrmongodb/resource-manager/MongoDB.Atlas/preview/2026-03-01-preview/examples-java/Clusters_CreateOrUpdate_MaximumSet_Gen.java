
import com.azure.resourcemanager.mongodbatlas.models.ClusterProperties;
import com.azure.resourcemanager.mongodbatlas.models.ClusterTier;

/**
 * Samples for Clusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Clusters_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Clusters_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        clustersCreateOrUpdateMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.clusters().define("myCluster").withExistingProject("rgopenapi", "myOrganization", "myProject")
            .withProperties(new ClusterProperties().withClusterTier(ClusterTier.FREE).withRegionName("eastus"))
            .create();
    }
}
