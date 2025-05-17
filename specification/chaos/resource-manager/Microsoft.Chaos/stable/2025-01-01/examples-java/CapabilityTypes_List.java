
/**
 * Samples for CapabilityTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CapabilityTypes_List.json
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
