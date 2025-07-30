
/**
 * Samples for ScheduledActionExtension ListByVms.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/ScheduledActionExtension_ListByVms_MaximumSet_Gen.json
     */
    /**
     * Sample code: ScheduledActionExtension_ListByVms_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void scheduledActionExtensionListByVmsMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.scheduledActionExtensions().listByVms("sazvpabfud", com.azure.core.util.Context.NONE);
    }
}
