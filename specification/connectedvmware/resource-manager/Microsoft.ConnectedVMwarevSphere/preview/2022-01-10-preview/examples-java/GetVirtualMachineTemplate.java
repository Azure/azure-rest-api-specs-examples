import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetVirtualMachineTemplate.json
     */
    /**
     * Sample code: GetVirtualMachineTemplate.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVirtualMachineTemplate(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachineTemplates().getByResourceGroupWithResponse("testrg", "WebFrontEndTemplate", Context.NONE);
    }
}
