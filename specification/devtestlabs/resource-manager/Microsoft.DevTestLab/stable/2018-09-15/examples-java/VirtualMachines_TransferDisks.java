/** Samples for VirtualMachines TransferDisks. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_TransferDisks.json
     */
    /**
     * Sample code: VirtualMachines_TransferDisks.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesTransferDisks(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .transferDisks("resourceGroupName", "{labName}", "{virtualmachineName}", com.azure.core.util.Context.NONE);
    }
}
