
/**
 * Samples for LogProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-03-01/getLogProfile.json
     */
    /**
     * Sample code: Get log profile.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getLogProfile(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getLogProfiles().getWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
