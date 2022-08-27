import com.azure.core.util.Context;

/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/stopVirtualMachine.json
     */
    /**
     * Sample code: stopVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void stopVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().stop("testrg123", "testlab", "template", Context.NONE);
    }
}
