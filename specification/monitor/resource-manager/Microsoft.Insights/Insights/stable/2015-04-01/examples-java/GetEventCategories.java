
/**
 * Samples for EventCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-04-01/GetEventCategories.json
     */
    /**
     * Sample code: Get event categories.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getEventCategories(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getEventCategories().list(com.azure.core.util.Context.NONE);
    }
}
