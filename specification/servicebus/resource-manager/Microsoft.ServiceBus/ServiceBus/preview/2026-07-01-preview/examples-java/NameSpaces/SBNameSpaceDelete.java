
/**
 * Samples for Namespaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceDelete.json
     */
    /**
     * Sample code: NameSpaceDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().delete("ArunMonocle", "sdk-Namespace-3285",
            com.azure.core.util.Context.NONE);
    }
}
