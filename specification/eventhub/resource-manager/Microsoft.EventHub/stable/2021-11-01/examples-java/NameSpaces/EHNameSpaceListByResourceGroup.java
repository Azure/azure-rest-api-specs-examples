
import com.azure.core.util.Context;

/** Samples for Namespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * EHNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: NamespaceListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespaceListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().listByResourceGroup("ResurceGroupSample",
            Context.NONE);
    }
}
