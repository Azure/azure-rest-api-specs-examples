import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachineTemplatesByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualMachineTemplatesByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualMachineTemplatesByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachineTemplates().listByResourceGroup("testrg", Context.NONE);
    }
}
