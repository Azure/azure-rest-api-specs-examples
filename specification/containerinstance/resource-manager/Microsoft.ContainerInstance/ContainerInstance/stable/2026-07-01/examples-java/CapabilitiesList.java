
/**
 * Samples for Location ListCapabilities.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CapabilitiesList.json
     */
    /**
     * Sample code: GetCapabilities.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void getCapabilities(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getLocations().listCapabilities("westus", com.azure.core.util.Context.NONE);
    }
}
