import com.azure.core.util.Context;

/** Samples for CloudLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/CloudLinks_Delete.json
     */
    /**
     * Sample code: CloudLinks_Delete.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.cloudLinks().delete("group1", "cloud1", "cloudLink1", Context.NONE);
    }
}
