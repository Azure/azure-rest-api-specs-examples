
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateEndpointConnections_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        privateEndpointConnectionsGetMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.privateEndpointConnections().getWithResponse("resourcegroupname", "elasticsanname",
            "privateendpointconnectionname", com.azure.core.util.Context.NONE);
    }
}
