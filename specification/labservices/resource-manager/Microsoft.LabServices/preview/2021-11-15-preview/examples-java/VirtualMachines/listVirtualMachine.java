import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByLab. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/VirtualMachines/listVirtualMachine.json
     */
    /**
     * Sample code: listVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().listByLab("testrg123", "testlab", null, Context.NONE);
    }
}
