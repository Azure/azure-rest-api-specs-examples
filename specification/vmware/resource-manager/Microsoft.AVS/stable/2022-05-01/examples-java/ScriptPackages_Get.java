import com.azure.core.util.Context;

/** Samples for ScriptPackages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptPackages_Get.json
     */
    /**
     * Sample code: ScriptPackages_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptPackagesGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptPackages().getWithResponse("group1", "{privateCloudName}", "{scriptPackageName}", Context.NONE);
    }
}
