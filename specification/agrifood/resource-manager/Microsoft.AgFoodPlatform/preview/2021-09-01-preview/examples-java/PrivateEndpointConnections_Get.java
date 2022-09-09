import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateEndpointConnections_Get.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Get.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void privateEndpointConnectionsGet(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "examples-rg", "examples-farmbeatsResourceName", "privateEndpointConnectionName", Context.NONE);
    }
}
