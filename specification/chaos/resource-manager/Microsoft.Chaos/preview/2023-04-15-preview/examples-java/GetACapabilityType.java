/** Samples for CapabilityTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetACapabilityType.json
     */
    /**
     * Sample code: Get a Capability Type for a virtual machine Target resource on westus2 location.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getACapabilityTypeForAVirtualMachineTargetResourceOnWestus2Location(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .capabilityTypes()
            .getWithResponse("westus2", "Microsoft-VirtualMachine", "Shutdown-1.0", com.azure.core.util.Context.NONE);
    }
}
