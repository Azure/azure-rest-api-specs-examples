
/**
 * Samples for Workspaces ListByOrganization.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Workspaces_ListByOrganization_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_ListByOrganization_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void workspacesListByOrganizationMaximumSetGen(
        com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.workspaces().listByOrganization("contoso-rg", "contoso-org", com.azure.core.util.Context.NONE);
    }
}
