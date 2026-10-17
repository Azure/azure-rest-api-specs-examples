
/**
 * Samples for EdgeMachineDisks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineDisks_Delete.json
     */
    /**
     * Sample code: EdgeMachineDisks_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void edgeMachineDisksDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineDisks().delete("test-rg", "EdgeMachine01", "disk-001", com.azure.core.util.Context.NONE);
    }
}
