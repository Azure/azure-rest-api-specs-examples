
/**
 * Samples for StandbyVirtualMachinePools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyVirtualMachinePools_ListBySubscription.json
     */
    /**
     * Sample code: StandbyVirtualMachinePools_ListBySubscription.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyVirtualMachinePoolsListBySubscription(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePools().list(com.azure.core.util.Context.NONE);
    }
}
