
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.NodeGroupSpec;
import com.azure.resourcemanager.mongocluster.models.NodeKind;
import java.util.Arrays;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mongocluster/DocumentDB.MongoCluster.Management/examples/2024-03-01-preview/
     * MongoClusters_PatchDiskSize.json
     */
    /**
     * Sample code: Updates the disk size on a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        updatesTheDiskSizeOnAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new MongoClusterUpdateProperties()
                .withNodeGroupSpecs(Arrays.asList(new NodeGroupSpec().withDiskSizeGB(256L).withKind(NodeKind.SHARD))))
            .apply();
    }
}
