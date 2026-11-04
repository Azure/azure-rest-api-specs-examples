
/**
 * Samples for Workspaces ListLinkedSaaS.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Workspaces_ListLinkedSaaS.json
     */
    /**
     * Sample code: Workspaces_ListLinkedSaaS.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void workspacesListLinkedSaaS(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.workspaces().listLinkedSaaSWithResponse("contoso-rg", "contoso-org", "contoso-prod",
            com.azure.core.util.Context.NONE);
    }
}
