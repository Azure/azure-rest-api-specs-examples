
/**
 * Samples for PrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/
     * PrivateEndPointConnectionList.json
     */
    /**
     * Sample code: PrivateEndPointConnectionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void privateEndPointConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getPrivateEndpointConnections().list("SDK-EventHub-4794",
            "sdk-Namespace-5828", com.azure.core.util.Context.NONE);
    }
}
