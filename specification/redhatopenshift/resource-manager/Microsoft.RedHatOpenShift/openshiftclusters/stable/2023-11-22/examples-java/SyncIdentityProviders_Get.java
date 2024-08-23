
/**
 * Samples for SyncIdentityProviders Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncIdentityProviders_Get.json
     */
    /**
     * Sample code: Gets a SyncIdentityProvider with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void getsASyncIdentityProviderWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncIdentityProviders().getWithResponse("resourceGroup", "resourceName", "childResourceName",
            com.azure.core.util.Context.NONE);
    }
}
