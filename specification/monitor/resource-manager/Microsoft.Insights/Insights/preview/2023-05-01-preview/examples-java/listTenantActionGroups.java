
/**
 * Samples for TenantActionGroups ListByManagementGroupId.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-05-01-preview/listTenantActionGroups.json
     */
    /**
     * Sample code: List tenant action groups at management group level.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listTenantActionGroupsAtManagementGroupLevel(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActionGroups().listByManagementGroupId("72f988bf-86f1-41af-91ab-2d7cd011db47",
            "72f988bf-86f1-41af-91ab-2d7cd011db47", com.azure.core.util.Context.NONE);
    }
}
