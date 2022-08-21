import com.azure.core.util.Context;

/** Samples for VirtualMachines AssessPatches. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineAssessPatches.json
     */
    /**
     * Sample code: Assess patch state of a machine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void assessPatchStateOfAMachine(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().assessPatches("myResourceGroupName", "myMachineName", Context.NONE);
    }
}
