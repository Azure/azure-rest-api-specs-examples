
import com.azure.resourcemanager.mongocluster.models.CreateMode;
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.MongoClusterRestoreParameters;
import java.time.OffsetDateTime;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-07-01/MongoClusters_CreatePITR.json
     */
    /**
     * Sample code: Creates a Mongo Cluster resource from a point in time restore.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void createsAMongoClusterResourceFromAPointInTimeRestore(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myMongoCluster").withRegion("westus2")
            .withExistingResourceGroup("TestResourceGroup")
            .withProperties(new MongoClusterProperties().withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
                .withRestoreParameters(new MongoClusterRestoreParameters()
                    .withPointInTimeUTC(OffsetDateTime.parse("2023-01-13T20:07:35Z")).withSourceResourceId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myOtherMongoCluster")))
            .create();
    }
}
