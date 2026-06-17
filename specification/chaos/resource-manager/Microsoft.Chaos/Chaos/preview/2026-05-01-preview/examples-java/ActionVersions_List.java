
/**
 * Samples for ActionVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ActionVersions_List.json
     */
    /**
     * Sample code: List all Action Versions for a given action.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllActionVersionsForAGivenAction(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.actionVersions().list("westus2", "microsoft-compute-shutdown", null, com.azure.core.util.Context.NONE);
    }
}
