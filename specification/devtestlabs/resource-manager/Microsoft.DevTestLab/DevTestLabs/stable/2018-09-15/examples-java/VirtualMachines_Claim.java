
/**
 * Samples for VirtualMachines Claim.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/VirtualMachines_Claim.json
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
