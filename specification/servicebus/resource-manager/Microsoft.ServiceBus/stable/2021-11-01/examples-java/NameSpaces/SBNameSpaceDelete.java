
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceDelete.json
     */
    /**
     * Sample code: NameSpaceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().delete("ArunMonocle",
            "sdk-Namespace-3285", com.azure.core.util.Context.NONE);
    }
}
