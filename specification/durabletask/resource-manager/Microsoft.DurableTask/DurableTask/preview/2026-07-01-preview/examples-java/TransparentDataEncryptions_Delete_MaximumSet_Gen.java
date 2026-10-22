
/**
 * Samples for TransparentDataEncryptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TransparentDataEncryptions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: TransparentDataEncryptions_Delete_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        transparentDataEncryptionsDeleteMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.transparentDataEncryptions().delete("rgdurabletask", "testscheduler", com.azure.core.util.Context.NONE);
    }
}
