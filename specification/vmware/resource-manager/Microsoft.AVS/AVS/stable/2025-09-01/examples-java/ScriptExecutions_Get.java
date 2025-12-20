
/**
 * Samples for ScriptExecutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptExecutions_Get.json
     */
    /**
     * Sample code: ScriptExecutions_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().getWithResponse("group1", "cloud1", "addSsoServer",
            com.azure.core.util.Context.NONE);
    }
}
