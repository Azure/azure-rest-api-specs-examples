import com.azure.resourcemanager.devcenter.models.EnableStatus;
import com.azure.resourcemanager.devcenter.models.ScheduledFrequency;
import com.azure.resourcemanager.devcenter.models.ScheduledType;

/** Samples for Schedules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Schedules_CreateDailyShutdownPoolSchedule.json
     */
    /**
     * Sample code: Schedules_CreateDailyShutdownPoolSchedule.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void schedulesCreateDailyShutdownPoolSchedule(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .schedules()
            .define("autoShutdown")
            .withExistingPool("rg1", "DevProject", "DevPool")
            .withTypePropertiesType(ScheduledType.STOP_DEV_BOX)
            .withFrequency(ScheduledFrequency.DAILY)
            .withTime("17:30")
            .withTimeZone("America/Los_Angeles")
            .withState(EnableStatus.ENABLED)
            .create();
    }
}
