
/**
 * Samples for LogProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-03-01/listLogProfile.json
     */
    /**
     * Sample code: List log profiles.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void listLogProfiles(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getLogProfiles().list(com.azure.core.util.Context.NONE);
    }
}
