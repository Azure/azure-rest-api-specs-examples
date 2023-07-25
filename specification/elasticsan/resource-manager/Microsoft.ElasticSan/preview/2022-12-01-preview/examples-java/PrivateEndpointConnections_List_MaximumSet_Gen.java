/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/PrivateEndpointConnections_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void privateEndpointConnectionsListMaximumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager
            .privateEndpointConnections()
            .list("resourcegroupname", "elasticsanname", com.azure.core.util.Context.NONE);
    }
}
