
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceGet.json
     */
    /**
     * Sample code: NameSpaceGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().getByResourceGroupWithResponse("ResurceGroupSample", "NamespaceSample",
            com.azure.core.util.Context.NONE);
    }
}
