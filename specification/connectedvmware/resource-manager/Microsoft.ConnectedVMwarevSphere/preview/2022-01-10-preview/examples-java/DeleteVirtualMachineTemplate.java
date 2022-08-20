import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteVirtualMachineTemplate.json
     */
    /**
     * Sample code: DeleteVirtualMachineTemplate.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVirtualMachineTemplate(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachineTemplates().delete("testrg", "WebFrontEndTemplate", null, Context.NONE);
    }
}
