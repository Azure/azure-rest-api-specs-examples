/** Samples for VirtualMachineTemplates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteVirtualMachineTemplate.json
     */
    /**
     * Sample code: DeleteVirtualMachineTemplate.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVirtualMachineTemplate(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .virtualMachineTemplates()
            .delete("testrg", "WebFrontEndTemplate", null, com.azure.core.util.Context.NONE);
    }
}
