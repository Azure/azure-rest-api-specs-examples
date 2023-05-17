/** Samples for Schedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Schedules_Delete.json
     */
    /**
     * Sample code: Schedules_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .schedules()
            .delete("rg1", "TestProject", "DevPool", "autoShutdown", null, com.azure.core.util.Context.NONE);
    }
}
