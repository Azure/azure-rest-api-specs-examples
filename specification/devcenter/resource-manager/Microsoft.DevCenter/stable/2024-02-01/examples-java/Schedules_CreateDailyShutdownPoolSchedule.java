
import com.azure.resourcemanager.devcenter.models.ScheduleEnableStatus;
import com.azure.resourcemanager.devcenter.models.ScheduledFrequency;
import com.azure.resourcemanager.devcenter.models.ScheduledType;

/**
 * Samples for Schedules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * Schedules_CreateDailyShutdownPoolSchedule.json
     */
    /**
     * Sample code: Schedules_CreateDailyShutdownPoolSchedule.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        schedulesCreateDailyShutdownPoolSchedule(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.schedules().define("autoShutdown").withExistingPool("rg1", "DevProject", "DevPool")
            .withTypePropertiesType(ScheduledType.STOP_DEV_BOX).withFrequency(ScheduledFrequency.DAILY)
            .withTime("17:30").withTimeZone("America/Los_Angeles").withState(ScheduleEnableStatus.ENABLED).create();
    }
}
