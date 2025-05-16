
/**
 * Samples for OperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OperationStatuses_Get.json
     */
    /**
     * Sample code: Gets Chaos Studio async operation status.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getsChaosStudioAsyncOperationStatus(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.operationStatuses().getWithResponse("westus2", "4bdadd97-207c-4de8-9bba-08339ae099c7",
            com.azure.core.util.Context.NONE);
    }
}
