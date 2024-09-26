
/**
 * Samples for StandbyContainerGroupPools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyContainerGroupPools_Get.json
     */
    /**
     * Sample code: StandbyContainerGroupPools_Get.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyContainerGroupPoolsGet(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPools().getByResourceGroupWithResponse("rgstandbypool", "pool",
            com.azure.core.util.Context.NONE);
    }
}
