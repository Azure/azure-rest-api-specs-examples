
/**
 * Samples for LogProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-03-01/deleteLogProfile.json
     */
    /**
     * Sample code: Delete log profile.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteLogProfile(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getLogProfiles().deleteWithResponse("Rac46PostSwapRG",
            com.azure.core.util.Context.NONE);
    }
}
