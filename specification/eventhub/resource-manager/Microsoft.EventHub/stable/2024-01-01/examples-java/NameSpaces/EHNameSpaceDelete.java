
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/
     * EHNameSpaceDelete.json
     */
    /**
     * Sample code: NameSpaceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().delete("ResurceGroupSample", "NamespaceSample",
            com.azure.core.util.Context.NONE);
    }
}
