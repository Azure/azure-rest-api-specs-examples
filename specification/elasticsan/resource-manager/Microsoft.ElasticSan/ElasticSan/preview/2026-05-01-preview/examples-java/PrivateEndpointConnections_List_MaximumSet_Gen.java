
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateEndpointConnections_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        privateEndpointConnectionsListMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.privateEndpointConnections().list("resourcegroupname", "elasticsanname",
            com.azure.core.util.Context.NONE);
    }
}
