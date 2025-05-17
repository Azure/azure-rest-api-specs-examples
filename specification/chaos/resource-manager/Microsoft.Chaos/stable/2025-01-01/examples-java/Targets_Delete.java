
/**
 * Samples for Targets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Targets_Delete.json
     */
    /**
     * Sample code: Delete a Target that extends a virtual machine resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        deleteATargetThatExtendsAVirtualMachineResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.targets().deleteWithResponse("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM",
            "Microsoft-Agent", com.azure.core.util.Context.NONE);
    }
}
