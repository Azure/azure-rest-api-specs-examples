import com.azure.core.util.Context;

/** Samples for ScriptCmdlets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptCmdlets_List.json
     */
    /**
     * Sample code: ScriptCmdlets_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptCmdletsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptCmdlets().list("group1", "{privateCloudName}", "{scriptPackageName}", Context.NONE);
    }
}
