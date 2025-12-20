
/**
 * Samples for ScriptExecutions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptExecutions_Delete.json
     */
    /**
     * Sample code: ScriptExecutions_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().delete("group1", "cloud1", "addSsoServer", com.azure.core.util.Context.NONE);
    }
}
