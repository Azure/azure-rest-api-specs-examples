import com.azure.core.util.Context;

/** Samples for Addons List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Addons_List.json
     */
    /**
     * Sample code: Addons_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void addonsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.addons().list("group1", "cloud1", Context.NONE);
    }
}
