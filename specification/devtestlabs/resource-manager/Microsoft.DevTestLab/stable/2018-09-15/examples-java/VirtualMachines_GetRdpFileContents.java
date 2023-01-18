/** Samples for VirtualMachines GetRdpFileContents. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_GetRdpFileContents.json
     */
    /**
     * Sample code: VirtualMachines_GetRdpFileContents.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualMachinesGetRdpFileContents(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualMachines()
            .getRdpFileContentsWithResponse(
                "resourceGroupName", "{labName}", "{vmName}", com.azure.core.util.Context.NONE);
    }
}
