/** Samples for CapabilityTypes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListCapabilityTypes.json
     */
    /**
     * Sample code: List all Capability Types for a virtual machine Target resource on westus2 location.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllCapabilityTypesForAVirtualMachineTargetResourceOnWestus2Location(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.capabilityTypes().list("westus2", "Microsoft-VirtualMachine", null, com.azure.core.util.Context.NONE);
    }
}
