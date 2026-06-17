
/**
 * Samples for Actions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Actions_List.json
     */
    /**
     * Sample code: List all Actions for westus2 location.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllActionsForWestus2Location(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.actions().list("westus2", null, com.azure.core.util.Context.NONE);
    }
}
