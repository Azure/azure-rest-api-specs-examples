import com.azure.core.util.Context;

/** Samples for CloudLinks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/CloudLinks_List.json
     */
    /**
     * Sample code: CloudLinks_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.cloudLinks().list("group1", "cloud1", Context.NONE);
    }
}
