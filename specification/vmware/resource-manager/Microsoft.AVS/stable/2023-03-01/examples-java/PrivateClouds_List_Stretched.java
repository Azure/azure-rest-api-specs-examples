/** Samples for PrivateClouds ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/PrivateClouds_List_Stretched.json
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
