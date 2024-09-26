
/**
 * Samples for StandbyVirtualMachinePools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyVirtualMachinePools_ListByResourceGroup.json
     */
    /**
     * Sample code: StandbyVirtualMachinePools_ListByResourceGroup.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyVirtualMachinePoolsListByResourceGroup(
        com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePools().listByResourceGroup("rgstandbypool", com.azure.core.util.Context.NONE);
    }
}
