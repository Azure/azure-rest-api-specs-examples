import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetVirtualMachineTemplate.json
     */
    /**
     * Sample code: GetVirtualMachineTemplate.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void getVirtualMachineTemplate(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .virtualMachineTemplates()
            .getByResourceGroupWithResponse("testrg", "HRVirtualMachineTemplate", Context.NONE);
    }
}
