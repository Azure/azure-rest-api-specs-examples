import com.azure.core.util.Context;
import com.azure.resourcemanager.devcenter.models.ScheduleUpdate;

/** Samples for Schedules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Schedules_Patch.json
     */
    /**
     * Sample code: Schedules_Update.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .schedules()
            .update(
                "rg1",
                "TestProject",
                "DevPool",
                "autoShutdown",
                new ScheduleUpdate().withTime("18:00"),
                null,
                Context.NONE);
    }
}
