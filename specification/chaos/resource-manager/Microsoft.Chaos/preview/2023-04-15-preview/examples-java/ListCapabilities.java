/** Samples for Capabilities List. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListCapabilities.json
     */
    /**
     * Sample code: List all Capabilities that extend a virtual machine Target resource.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllCapabilitiesThatExtendAVirtualMachineTargetResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .capabilities()
            .list(
                "exampleRG",
                "Microsoft.Compute",
                "virtualMachines",
                "exampleVM",
                "Microsoft-VirtualMachine",
                null,
                com.azure.core.util.Context.NONE);
    }
}
