
/**
 * Samples for Targets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Targets_List.json
     */
    /**
     * Sample code: List all Targets that extend a virtual machine resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        listAllTargetsThatExtendAVirtualMachineResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.targets().list("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", null,
            com.azure.core.util.Context.NONE);
    }
}
