import com.azure.core.util.Context;

/** Samples for Datastores Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Datastores_Get.json
     */
    /**
     * Sample code: Datastores_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void datastoresGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.datastores().getWithResponse("group1", "cloud1", "cluster1", "datastore1", Context.NONE);
    }
}
