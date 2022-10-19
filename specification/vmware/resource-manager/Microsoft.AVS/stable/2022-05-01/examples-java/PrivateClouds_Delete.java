import com.azure.core.util.Context;

/** Samples for PrivateClouds Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PrivateClouds_Delete.json
     */
    /**
     * Sample code: PrivateClouds_Delete.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().delete("group1", "cloud1", Context.NONE);
    }
}
