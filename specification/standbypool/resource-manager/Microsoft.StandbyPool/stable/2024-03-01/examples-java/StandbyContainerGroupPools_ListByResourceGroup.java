
/**
 * Samples for StandbyContainerGroupPools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyContainerGroupPools_ListByResourceGroup.json
     */
    /**
     * Sample code: StandbyContainerGroupPools_ListByResourceGroup.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyContainerGroupPoolsListByResourceGroup(
        com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPools().listByResourceGroup("rgstandbypool", com.azure.core.util.Context.NONE);
    }
}
