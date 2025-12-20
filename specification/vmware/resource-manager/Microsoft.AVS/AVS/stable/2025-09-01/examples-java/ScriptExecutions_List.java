
/**
 * Samples for ScriptExecutions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptExecutions_List.json
     */
    /**
     * Sample code: ScriptExecutions_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().list("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
