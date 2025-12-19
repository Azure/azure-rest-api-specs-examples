
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/VirtualMachines_List.json
     */
    /**
     * Sample code: VirtualMachines_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void virtualMachinesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.virtualMachines().list("group1", "cloud1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
