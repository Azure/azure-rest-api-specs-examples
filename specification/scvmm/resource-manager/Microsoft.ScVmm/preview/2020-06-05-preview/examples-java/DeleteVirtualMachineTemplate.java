import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVirtualMachineTemplate.json
     */
    /**
     * Sample code: DeleteVirtualMachineTemplate.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteVirtualMachineTemplate(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().delete("testrg", "HRVirtualMachineTemplate", null, Context.NONE);
    }
}
