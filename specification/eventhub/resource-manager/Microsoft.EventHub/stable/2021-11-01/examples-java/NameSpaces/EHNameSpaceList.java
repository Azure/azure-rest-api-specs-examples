
import com.azure.core.util.Context;

/** Samples for Namespaces List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceList.
     * json
     */
    /**
     * Sample code: NamespacesListBySubscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().list(Context.NONE);
    }
}
