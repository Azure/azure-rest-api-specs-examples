import com.azure.core.util.Context;

/** Samples for MachineExtensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/LISTExtension.json
     */
    /**
     * Sample code: Get all Machine Extensions.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getAllMachineExtensions(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.machineExtensions().list("myResourceGroup", "myMachine", null, Context.NONE);
    }
}
