
/**
 * Samples for StandbyVirtualMachinePools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/StandbyVirtualMachinePools_Get.json
     */
    /**
     * Sample code: StandbyVirtualMachinePools_Get.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyVirtualMachinePoolsGet(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePools().getByResourceGroupWithResponse("rgstandbypool", "pool",
            com.azure.core.util.Context.NONE);
    }
}
