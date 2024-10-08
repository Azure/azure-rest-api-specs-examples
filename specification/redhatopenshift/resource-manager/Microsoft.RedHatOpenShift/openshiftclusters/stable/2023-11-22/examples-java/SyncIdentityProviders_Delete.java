
/**
 * Samples for SyncIdentityProviders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncIdentityProviders_Delete.json
     */
    /**
     * Sample code: Deletes a SyncIdentityProvider with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void deletesASyncIdentityProviderWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncIdentityProviders().deleteWithResponse("resourceGroup", "resourceName", "childResourceName",
            com.azure.core.util.Context.NONE);
    }
}
