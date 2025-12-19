
/**
 * Samples for PrivateClouds ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_List_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_List_Stretched.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().listByResourceGroup("group1", com.azure.core.util.Context.NONE);
    }
}
