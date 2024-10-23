
import com.azure.resourcemanager.elasticsan.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/
     * PrivateEndpointConnections_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Create_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        privateEndpointConnectionsCreateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.privateEndpointConnections().define("privateendpointconnectionname")
            .withExistingElasticSan("resourcegroupname", "elasticsanname")
            .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()).create();
    }
}
