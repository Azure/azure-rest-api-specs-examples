
/**
 * Samples for TenantActionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/getTenantActionGroup.json
     */
    /**
     * Sample code: Get an action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getAnActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActionGroups().getWithResponse("72f988bf-86f1-41af-91ab-2d7cd011db47",
            "testTenantActionGroup", "72f988bf-86f1-41af-91ab-2d7cd011db47", com.azure.core.util.Context.NONE);
    }
}
