
/**
 * Samples for PrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * PrivateEndPointConnectionGet.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpacePrivateEndPointConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getPrivateEndpointConnections().getWithResponse(
            "SDK-ServiceBus-4794", "sdk-Namespace-5828", "privateEndpointConnectionName",
            com.azure.core.util.Context.NONE);
    }
}
