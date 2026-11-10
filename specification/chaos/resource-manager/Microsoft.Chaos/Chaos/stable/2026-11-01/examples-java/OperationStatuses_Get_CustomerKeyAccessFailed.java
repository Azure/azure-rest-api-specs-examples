
/**
 * Samples for OperationStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/OperationStatuses_Get_CustomerKeyAccessFailed.json
     */
    /**
     * Sample code: Get a failed Workspace encryption operation status.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getAFailedWorkspaceEncryptionOperationStatus(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.operationStatuses().getWithResponse("westus2", "4bdadd97-207c-4de8-9bba-08339ae099c7",
            com.azure.core.util.Context.NONE);
    }
}
