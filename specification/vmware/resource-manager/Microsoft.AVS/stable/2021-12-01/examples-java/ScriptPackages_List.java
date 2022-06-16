import com.azure.core.util.Context;

/** Samples for ScriptPackages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptPackages_List.json
     */
    /**
     * Sample code: ScriptPackages_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptPackagesList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptPackages().list("group1", "{privateCloudName}", Context.NONE);
    }
}
