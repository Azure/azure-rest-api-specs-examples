
import com.azure.core.util.Context;

/** Samples for Namespaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: NameSpaceListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().listByResourceGroup("ArunMonocle",
            Context.NONE);
    }
}
