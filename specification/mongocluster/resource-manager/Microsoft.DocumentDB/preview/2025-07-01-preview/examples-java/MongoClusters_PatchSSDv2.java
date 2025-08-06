
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.StorageProperties;
import com.azure.resourcemanager.mongocluster.models.StorageType;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_PatchSSDv2.json
     */
    /**
     * Sample code: Updates the Premium SSDv2 size, IOPS and throughput on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void updatesThePremiumSSDv2SizeIOPSAndThroughputOnAMongoClusterResource(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties().withStorage(new StorageProperties()
            .withSizeGb(128L).withType(StorageType.PREMIUM_SSDV2).withIops(5000L).withThroughput(1000L))).apply();
    }
}
