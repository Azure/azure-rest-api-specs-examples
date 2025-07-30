
/**
 * Samples for OccurrenceExtension ListOccurrenceByVms.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-15-preview/OccurrenceExtension_ListOccurrenceByVms_MaximumSet_Gen.json
     */
    /**
     * Sample code: OccurrenceExtension_ListOccurrenceByVms_MaximumSet.
     * 
     * @param manager Entry point to ComputeScheduleManager.
     */
    public static void occurrenceExtensionListOccurrenceByVmsMaximumSet(
        com.azure.resourcemanager.computeschedule.ComputeScheduleManager manager) {
        manager.occurrenceExtensions().listOccurrenceByVms("sazvpabfud", com.azure.core.util.Context.NONE);
    }
}
