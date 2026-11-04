
/**
 * Samples for Organizations ListWorkspaces.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_ListWorkspaces.json
     */
    /**
     * Sample code: Organizations_ListWorkspaces.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        organizationsListWorkspaces(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().listWorkspacesWithResponse("contoso-rg", "contoso-org",
            com.azure.core.util.Context.NONE);
    }
}
