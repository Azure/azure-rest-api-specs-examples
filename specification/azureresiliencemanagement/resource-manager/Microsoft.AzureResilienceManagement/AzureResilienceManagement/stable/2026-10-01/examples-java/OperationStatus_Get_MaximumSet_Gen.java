
/**
 * Samples for OperationStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/OperationStatus_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: OperationStatus_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void operationStatusGetMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.operationStatus().getWithResponse("eastus", "12345678-1234-1234-1234-123456789012",
            com.azure.core.util.Context.NONE);
    }
}
