
/**
 * Samples for IscsiPaths Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_Delete.json
     */
    /**
     * Sample code: IscsiPaths_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void iscsiPathsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.iscsiPaths().delete("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
