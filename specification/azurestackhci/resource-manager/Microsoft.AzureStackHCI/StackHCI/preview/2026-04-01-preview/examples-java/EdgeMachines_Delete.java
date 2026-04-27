
/**
 * Samples for EdgeMachines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachines_Delete.json
     */
    /**
     * Sample code: EdgeMachines_Delete_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachinesDeleteMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().delete("ArcInstance-rg", "machine-1", com.azure.core.util.Context.NONE);
    }
}
