
/**
 * Samples for VirtualMachines ListApplicableSchedules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_ListApplicableSchedules.json
     */
    /**
     * Sample code: VirtualMachines_ListApplicableSchedules.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void
        virtualMachinesListApplicableSchedules(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().listApplicableSchedulesWithResponse("resourceGroupName", "{labName}", "{vmName}",
            com.azure.core.util.Context.NONE);
    }
}
