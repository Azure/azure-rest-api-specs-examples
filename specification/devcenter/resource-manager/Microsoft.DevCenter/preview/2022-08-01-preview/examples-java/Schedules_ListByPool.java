import com.azure.core.util.Context;

/** Samples for Schedules ListByPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_ListByPool.json
     */
    /**
     * Sample code: Schedules_ListByPool.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesListByPool(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.schedules().listByPool("rg1", "TestProject", "DevPool", null, Context.NONE);
    }
}
