
/**
 * Samples for PrivateClouds List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_ListInSubscription_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_ListInSubscription_Stretched.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListInSubscriptionStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().list(com.azure.core.util.Context.NONE);
    }
}
