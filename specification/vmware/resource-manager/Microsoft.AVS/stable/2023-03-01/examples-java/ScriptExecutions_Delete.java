/** Samples for ScriptExecutions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/ScriptExecutions_Delete.json
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
