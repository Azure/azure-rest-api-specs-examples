
/**
 * Samples for ScriptCmdlets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptCmdlets_List.json
     */
    /**
     * Sample code: ScriptCmdlets_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptCmdletsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptCmdlets().list("group1", "cloud1", "package@1.0.2", com.azure.core.util.Context.NONE);
    }
}
