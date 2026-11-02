
/**
 * Samples for SystemReadinessOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/SystemReadinessOperations_Get.json
     */
    /**
     * Sample code: SystemReadinessOperations_Get.
     * 
     * @param manager Entry point to EdgeOperatorManager.
     */
    public static void
        systemReadinessOperationsGet(com.azure.resourcemanager.edgeoperator.EdgeOperatorManager manager) {
        manager.systemReadinessOperations().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
