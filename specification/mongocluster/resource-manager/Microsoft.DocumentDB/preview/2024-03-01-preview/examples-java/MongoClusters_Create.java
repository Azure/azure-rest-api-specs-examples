
import com.azure.resourcemanager.mongocluster.models.MongoClusterProperties;
import com.azure.resourcemanager.mongocluster.models.NodeGroupSpec;
import com.azure.resourcemanager.mongocluster.models.NodeKind;
import java.util.Arrays;

/**
 * Samples for MongoClusters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-01-preview/MongoClusters_Create.json
     */
    /**
     * Sample code: Creates a new Mongo Cluster resource.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void
        createsANewMongoClusterResource(com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        manager.mongoClusters().define("myMongoCluster").withRegion("westus2")
            .withExistingResourceGroup(
                "TestResourceGroup")
            .withProperties(new MongoClusterProperties().withAdministratorLogin("mongoAdmin")
                .withAdministratorLoginPassword("fakeTokenPlaceholder").withServerVersion("5.0")
                .withNodeGroupSpecs(Arrays.asList(new NodeGroupSpec().withSku("M30").withDiskSizeGB(128L)
                    .withEnableHa(true).withKind(NodeKind.SHARD).withNodeCount(1))))
            .create();
    }
}
