
/**
 * Samples for PrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * PrivateEndPointConnectionDelete.json
     */
    /**
     * Sample code: NameSpacePrivateEndPointConnectionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpacePrivateEndPointConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getPrivateEndpointConnections().delete("ArunMonocle",
            "sdk-Namespace-3285", "928c44d5-b7c6-423b-b6fa-811e0c27b3e0", com.azure.core.util.Context.NONE);
    }
}
