
/**
 * Samples for Workspaces GetDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Workspaces_GetDetails.json
     */
    /**
     * Sample code: Workspaces_GetDetails.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void workspacesGetDetails(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.workspaces().getDetailsWithResponse("contoso-rg", "contoso-org", "contoso-prod",
            com.azure.core.util.Context.NONE);
    }
}
