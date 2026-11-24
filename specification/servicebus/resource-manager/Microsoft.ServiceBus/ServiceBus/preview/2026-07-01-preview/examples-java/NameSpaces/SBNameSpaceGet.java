
/**
 * Samples for Namespaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceGet.json
     */
    /**
     * Sample code: NameSpaceGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().getByResourceGroupWithResponse("ArunMonocle", "sdk-Namespace-2924",
            com.azure.core.util.Context.NONE);
    }
}
