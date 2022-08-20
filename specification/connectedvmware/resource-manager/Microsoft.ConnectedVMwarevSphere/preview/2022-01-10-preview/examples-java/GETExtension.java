import com.azure.core.util.Context;

/** Samples for MachineExtensions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GETExtension.json
     */
    /**
     * Sample code: Get Machine Extension.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getMachineExtension(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .machineExtensions()
            .getWithResponse("myResourceGroup", "myMachine", "CustomScriptExtension", Context.NONE);
    }
}
