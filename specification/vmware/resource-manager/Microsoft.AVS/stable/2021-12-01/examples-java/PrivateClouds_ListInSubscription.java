import com.azure.core.util.Context;

/** Samples for PrivateClouds List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_ListInSubscription.json
     */
    /**
     * Sample code: PrivateClouds_ListInSubscription.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListInSubscription(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().list(Context.NONE);
    }
}
