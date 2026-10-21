
/**
 * Samples for TransparentDataEncryptions ListByScheduler.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TransparentDataEncryptions_ListByScheduler_MaximumSet_Gen.json
     */
    /**
     * Sample code: TransparentDataEncryptions_ListByScheduler_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void transparentDataEncryptionsListBySchedulerMaximumSet(
        com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.transparentDataEncryptions().listByScheduler("rgdurabletask", "testscheduler",
            com.azure.core.util.Context.NONE);
    }
}
