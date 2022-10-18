import com.azure.core.util.Context;

/** Samples for ScriptCmdlets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptCmdlets_Get.json
     */
    /**
     * Sample code: ScriptCmdlets_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptCmdletsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .scriptCmdlets()
            .getWithResponse(
                "group1", "{privateCloudName}", "{scriptPackageName}", "New-ExternalSsoDomain", Context.NONE);
    }
}
