
/**
 * Samples for ActionGroups ListNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForActionGroups_List.json
     */
    /**
     * Sample code: Get an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().listNSP("exampleRG", "someActionGroup",
            com.azure.core.util.Context.NONE);
    }
}
