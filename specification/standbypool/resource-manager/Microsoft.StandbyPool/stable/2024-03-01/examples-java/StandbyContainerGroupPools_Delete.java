
/**
 * Samples for StandbyContainerGroupPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyContainerGroupPools_Delete.json
     */
    /**
     * Sample code: StandbyContainerGroupPools_Delete.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyContainerGroupPoolsDelete(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPools().delete("rgstandbypool", "pool", com.azure.core.util.Context.NONE);
    }
}
