
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/PrivateLinkResources_List
     * .json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.privateLinkResources().list("myResourceGroup", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
