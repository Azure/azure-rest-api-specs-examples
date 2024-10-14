
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import com.azure.resourcemanager.mongocluster.models.NodeGroupSpec;
import com.azure.resourcemanager.mongocluster.models.NodeKind;
import com.azure.resourcemanager.mongocluster.models.PublicNetworkAccess;
import java.util.Arrays;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_Update.json
     */
    /**
     * Sample code: Updates a Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        updatesAMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties().withAdministratorLogin("mongoAdmin")
            .withAdministratorLoginPassword("fakeTokenPlaceholder").withServerVersion("5.0")
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED).withNodeGroupSpecs(Arrays.asList(new NodeGroupSpec()
                .withSku("M50").withDiskSizeGB(256L).withEnableHa(true).withKind(NodeKind.SHARD).withNodeCount(1))))
            .apply();
    }
}
