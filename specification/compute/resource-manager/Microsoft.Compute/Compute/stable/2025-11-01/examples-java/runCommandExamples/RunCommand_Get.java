
/**
 * Samples for VirtualMachineRunCommands Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/runCommandExamples/RunCommand_Get.json
     */
    /**
     * Sample code: VirtualMachineRunCommandGet.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineRunCommandGet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineRunCommands().getWithResponse("SoutheastAsia", "RunPowerShellScript",
            com.azure.core.util.Context.NONE);
    }
}
