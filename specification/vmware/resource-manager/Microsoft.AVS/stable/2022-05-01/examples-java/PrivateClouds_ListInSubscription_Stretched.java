import com.azure.core.util.Context;

/** Samples for PrivateClouds List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PrivateClouds_ListInSubscription_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_ListInSubscription_Stretched.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListInSubscriptionStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().list(Context.NONE);
    }
}
