import com.azure.core.util.Context;

/** Samples for VirtualMachines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/VirtualMachines/getVirtualMachine.json
     */
    /**
     * Sample code: getVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().getWithResponse("testrg123", "testlab", "template", Context.NONE);
    }
}
