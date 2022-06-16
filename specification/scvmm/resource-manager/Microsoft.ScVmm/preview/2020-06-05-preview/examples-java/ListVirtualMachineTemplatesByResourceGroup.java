import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualMachineTemplatesByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualMachineTemplatesByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVirtualMachineTemplatesByResourceGroup(
        com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().listByResourceGroup("testrg", Context.NONE);
    }
}
