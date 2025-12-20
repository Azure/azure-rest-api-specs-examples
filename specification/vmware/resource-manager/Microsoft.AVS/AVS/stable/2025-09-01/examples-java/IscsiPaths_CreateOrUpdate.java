
import com.azure.resourcemanager.avs.fluent.models.IscsiPathInner;

/**
 * Samples for IscsiPaths CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/IscsiPaths_CreateOrUpdate.json
     */
    /**
     * Sample code: IscsiPaths_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void iscsiPathsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.iscsiPaths().createOrUpdate("group1", "cloud1", new IscsiPathInner().withNetworkBlock("192.168.0.0/24"),
            com.azure.core.util.Context.NONE);
    }
}
