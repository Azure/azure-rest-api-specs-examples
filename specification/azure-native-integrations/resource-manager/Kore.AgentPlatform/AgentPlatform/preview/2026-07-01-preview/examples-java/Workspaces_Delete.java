
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Workspaces_Delete.json
     */
    /**
     * Sample code: Workspaces_Delete.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void workspacesDelete(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.workspaces().delete("contoso-rg", "contoso-org", "contoso-prod", com.azure.core.util.Context.NONE);
    }
}
