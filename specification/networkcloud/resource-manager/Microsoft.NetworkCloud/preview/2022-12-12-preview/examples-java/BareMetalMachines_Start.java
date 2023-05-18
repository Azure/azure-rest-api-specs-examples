/** Samples for BareMetalMachines Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachines_Start.json
     */
    /**
     * Sample code: Start bare metal machine.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void startBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bareMetalMachines()
            .start("resourceGroupName", "bareMetalMachineName", com.azure.core.util.Context.NONE);
    }
}
