import com.azure.core.util.Context;

/** Samples for CloudLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/CloudLinks_Get.json
     */
    /**
     * Sample code: CloudLinks_Get.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.cloudLinks().getWithResponse("group1", "cloud1", "cloudLink1", Context.NONE);
    }
}
