
import com.azure.resourcemanager.devcenter.models.Schedule;

/**
 * Samples for Schedules Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Schedules_Patch.json
     */
    /**
     * Sample code: Schedules_Update.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        Schedule resource = manager.schedules()
            .getWithResponse("rg1", "TestProject", "DevPool", "autoShutdown", null, com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTime("18:00").apply();
    }
}
