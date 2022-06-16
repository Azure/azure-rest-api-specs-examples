import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualMachineTemplatesBySubscription.json
     */
    /**
     * Sample code: ListVirtualMachineTemplatesBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVirtualMachineTemplatesBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().list(Context.NONE);
    }
}
