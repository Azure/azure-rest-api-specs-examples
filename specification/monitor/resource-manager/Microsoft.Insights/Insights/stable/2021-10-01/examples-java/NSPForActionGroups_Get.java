
/**
 * Samples for ActionGroups GetNSP.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-10-01/NSPForActionGroups_Get.json
     */
    /**
     * Sample code: Get NSP config by name for an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getNSPConfigByNameForAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActionGroups().getNSPWithResponse("exampleRG", "someActionGroup",
            "somePerimeterConfiguration", com.azure.core.util.Context.NONE);
    }
}
