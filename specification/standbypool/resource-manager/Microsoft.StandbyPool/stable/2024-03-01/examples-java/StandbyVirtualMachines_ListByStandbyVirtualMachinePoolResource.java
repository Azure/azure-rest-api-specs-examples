
/**
 * Samples for StandbyVirtualMachines ListByStandbyVirtualMachinePoolResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyVirtualMachines_ListByStandbyVirtualMachinePoolResource.json
     */
    /**
     * Sample code: StandbyVirtualMachines_ListByStandbyVirtualMachinePoolResource.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyVirtualMachinesListByStandbyVirtualMachinePoolResource(
        com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachines().listByStandbyVirtualMachinePoolResource("rgstandbypool", "pool",
            com.azure.core.util.Context.NONE);
    }
}
