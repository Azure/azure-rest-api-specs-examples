/** Samples for ScriptCmdlets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/ScriptCmdlets_Get.json
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
                "group1", "cloud1", "package@1.0.2", "New-ExternalSsoDomain", com.azure.core.util.Context.NONE);
    }
}
