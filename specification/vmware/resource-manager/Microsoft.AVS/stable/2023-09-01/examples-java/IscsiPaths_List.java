
/**
 * Samples for IscsiPaths ListByPrivateCloud.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/IscsiPaths_List.json
     */
    /**
     * Sample code: IscsiPaths_ListByPrivateCloud.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void iscsiPathsListByPrivateCloud(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.iscsiPaths().listByPrivateCloud("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
