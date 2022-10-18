import com.azure.core.util.Context;

/** Samples for ScriptExecutions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptExecutions_List.json
     */
    /**
     * Sample code: ScriptExecutions_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().list("group1", "{privateCloudName}", Context.NONE);
    }
}
