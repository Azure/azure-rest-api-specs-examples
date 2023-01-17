/** Samples for VirtualMachines Claim. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_Claim.json
     */
    /**
     * Sample code: VirtualMachines_Claim.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesClaim(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.virtualMachines().claim("resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
