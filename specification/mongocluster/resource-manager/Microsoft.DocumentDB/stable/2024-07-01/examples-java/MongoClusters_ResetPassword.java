
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_ResetPassword.json
     */
    /**
     * Sample code: Resets the administrator login password.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        resetsTheAdministratorLoginPassword(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties().withAdministratorLogin("mongoAdmin")
            .withAdministratorLoginPassword("fakeTokenPlaceholder")).apply();
    }
}
