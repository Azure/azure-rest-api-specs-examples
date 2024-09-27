
/**
 * Samples for StandbyVirtualMachinePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyVirtualMachinePools_Delete.json
     */
    /**
     * Sample code: StandbyVirtualMachinePools_Delete.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyVirtualMachinePoolsDelete(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePools().delete("rgstandbypool", "pool", com.azure.core.util.Context.NONE);
    }
}
