/** Samples for BareMetalMachines Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_Restart.json
     */
    /**
     * Sample code: Restart bare metal machine.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void restartBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bareMetalMachines()
            .restart("resourceGroupName", "bareMetalMachineName", com.azure.core.util.Context.NONE);
    }
}
