
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceList.json
     */
    /**
     * Sample code: NameSpaceList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().list(com.azure.core.util.Context.NONE);
    }
}
