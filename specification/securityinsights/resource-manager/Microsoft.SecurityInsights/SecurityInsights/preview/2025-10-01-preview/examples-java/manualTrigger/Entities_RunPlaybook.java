
/**
 * Samples for Entities RunPlaybook.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/manualTrigger/Entities_RunPlaybook.json
     */
    /**
     * Sample code: Entities_RunPlaybook.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void entitiesRunPlaybook(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().runPlaybookWithResponse("myRg", "myWorkspace", "72e01a22-5cd2-4139-a149-9f2736ff2ar2", null,
            com.azure.core.util.Context.NONE);
    }
}
