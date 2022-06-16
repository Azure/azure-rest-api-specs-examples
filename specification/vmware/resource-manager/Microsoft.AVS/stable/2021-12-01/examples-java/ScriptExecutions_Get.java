import com.azure.core.util.Context;

/** Samples for ScriptExecutions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_Get.json
     */
    /**
     * Sample code: ScriptExecutions_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().getWithResponse("group1", "cloud1", "addSsoServer", Context.NONE);
    }
}
