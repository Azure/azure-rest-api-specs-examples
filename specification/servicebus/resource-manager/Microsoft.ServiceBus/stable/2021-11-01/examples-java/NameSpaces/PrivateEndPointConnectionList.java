
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * PrivateEndPointConnectionList.json
     */
    /**
     * Sample code: NameSpaceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getPrivateEndpointConnections()
            .list("SDK-ServiceBus-4794", "sdk-Namespace-5828", com.azure.core.util.Context.NONE);
    }
}
