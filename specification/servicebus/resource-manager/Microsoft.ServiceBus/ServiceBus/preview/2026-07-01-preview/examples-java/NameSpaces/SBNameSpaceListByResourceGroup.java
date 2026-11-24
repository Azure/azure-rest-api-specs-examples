
/**
 * Samples for Namespaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceListByResourceGroup.json
     */
    /**
     * Sample code: NameSpaceListByResourceGroup.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceListByResourceGroup(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().listByResourceGroup("ArunMonocle", com.azure.core.util.Context.NONE);
    }
}
