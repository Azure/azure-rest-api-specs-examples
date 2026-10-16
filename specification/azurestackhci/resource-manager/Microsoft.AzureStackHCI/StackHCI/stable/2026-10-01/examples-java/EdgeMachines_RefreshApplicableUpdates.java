
/**
 * Samples for EdgeMachines RefreshApplicableUpdates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachines_RefreshApplicableUpdates.json
     */
    /**
     * Sample code: Refresh edge machine applicable updates.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        refreshEdgeMachineApplicableUpdates(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().refreshApplicableUpdates("test-rg", "testEdgeMachine", com.azure.core.util.Context.NONE);
    }
}
