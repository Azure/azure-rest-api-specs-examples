
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceList.json
     */
    /**
     * Sample code: NameSpaceList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().list(com.azure.core.util.Context.NONE);
    }
}
