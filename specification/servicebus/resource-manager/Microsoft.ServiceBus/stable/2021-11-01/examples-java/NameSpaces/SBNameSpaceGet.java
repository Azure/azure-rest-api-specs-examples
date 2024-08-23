
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceGet.json
     */
    /**
     * Sample code: NameSpaceGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces()
            .getByResourceGroupWithResponse("ArunMonocle", "sdk-Namespace-2924", com.azure.core.util.Context.NONE);
    }
}
