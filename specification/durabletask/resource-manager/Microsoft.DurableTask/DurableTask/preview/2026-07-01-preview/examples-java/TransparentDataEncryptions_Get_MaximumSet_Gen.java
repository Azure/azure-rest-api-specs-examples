
/**
 * Samples for TransparentDataEncryptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TransparentDataEncryptions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: TransparentDataEncryptions_Get_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void
        transparentDataEncryptionsGetMaximumSet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.transparentDataEncryptions().getWithResponse("rgdurabletask", "testscheduler",
            com.azure.core.util.Context.NONE);
    }
}
