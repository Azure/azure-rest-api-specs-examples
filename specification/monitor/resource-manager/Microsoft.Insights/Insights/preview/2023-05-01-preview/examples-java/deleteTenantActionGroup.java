
/**
 * Samples for TenantActionGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/deleteTenantActionGroup.json
     */
    /**
     * Sample code: Delete a tenant action group.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deleteATenantActionGroup(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActionGroups().deleteWithResponse("72f988bf-86f1-41af-91ab-2d7cd011db47",
            "testTenantActionGroup", "72f988bf-86f1-41af-91ab-2d7cd011db47", com.azure.core.util.Context.NONE);
    }
}
