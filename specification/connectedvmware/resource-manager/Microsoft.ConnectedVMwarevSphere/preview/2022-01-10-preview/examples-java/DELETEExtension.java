import com.azure.core.util.Context;

/** Samples for MachineExtensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DELETEExtension.json
     */
    /**
     * Sample code: Delete a Machine Extension.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteAMachineExtension(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.machineExtensions().delete("myResourceGroup", "myMachine", "MMA", Context.NONE);
    }
}
