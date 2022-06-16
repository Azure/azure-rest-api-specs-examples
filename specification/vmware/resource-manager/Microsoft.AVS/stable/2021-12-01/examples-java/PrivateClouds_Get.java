import com.azure.core.util.Context;

/** Samples for PrivateClouds GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_Get.json
     */
    /**
     * Sample code: PrivateClouds_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().getByResourceGroupWithResponse("group1", "cloud1", Context.NONE);
    }
}
