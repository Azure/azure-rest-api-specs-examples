
/**
 * Samples for Workspaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Workspaces_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        workspacesGetMaximumSetGen(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.workspaces().getWithResponse("contoso-rg", "contoso-org", "contoso-prod",
            com.azure.core.util.Context.NONE);
    }
}
