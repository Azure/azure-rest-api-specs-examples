
/**
 * Samples for VirtualMachines RetrieveBootDiagnosticsData.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_RetrieveBootDiagnosticsData.json
     */
    /**
     * Sample code: RetrieveBootDiagnosticsData of a virtual machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        retrieveBootDiagnosticsDataOfAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().retrieveBootDiagnosticsDataWithResponse("ResourceGroup", "VMName",
            60, com.azure.core.util.Context.NONE);
    }
}
