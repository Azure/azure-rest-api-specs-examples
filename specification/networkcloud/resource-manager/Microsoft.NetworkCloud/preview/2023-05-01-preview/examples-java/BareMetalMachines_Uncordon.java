/** Samples for BareMetalMachines Uncordon. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_Uncordon.json
     */
    /**
     * Sample code: Uncordon bare metal machine.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void uncordonBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bareMetalMachines()
            .uncordon("resourceGroupName", "bareMetalMachineName", com.azure.core.util.Context.NONE);
    }
}
