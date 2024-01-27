
import com.azure.core.util.Context;

/** Samples for Namespaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceGet.
     * json
     */
    /**
     * Sample code: NameSpaceGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().getByResourceGroupWithResponse("ResurceGroupSample",
            "NamespaceSample", Context.NONE);
    }
}
