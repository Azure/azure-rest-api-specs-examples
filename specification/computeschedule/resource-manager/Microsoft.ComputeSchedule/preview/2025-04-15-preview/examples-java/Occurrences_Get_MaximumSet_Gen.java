
/**
 * Samples for Occurrences Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/Occurrences_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Occurrences_Get_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void
        occurrencesGetMaximumSet(com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.occurrences().getWithResponse("rgcomputeschedule", "myScheduledAction",
            "67b5bada-4772-43fc-8dbb-402476d98a45", com.azure.core.util.Context.NONE);
    }
}
