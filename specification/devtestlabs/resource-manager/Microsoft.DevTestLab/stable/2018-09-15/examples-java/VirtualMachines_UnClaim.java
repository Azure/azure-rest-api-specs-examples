/** Samples for VirtualMachines UnClaim. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_UnClaim.json
     */
    /**
     * Sample code: VirtualMachines_UnClaim.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesUnClaim(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .unClaim("resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
