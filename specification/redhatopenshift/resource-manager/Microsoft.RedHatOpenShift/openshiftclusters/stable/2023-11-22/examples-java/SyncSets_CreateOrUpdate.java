
/**
 * Samples for SyncSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/SyncSets_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a SyncSet with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void createsOrUpdatesASyncSetWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.syncSets().define("childResourceName").withExistingOpenshiftcluster("resourceGroup", "resourceName")
            .withResources(
                "eyAKICAiYXBpVmVyc2lvbiI6ICJoaXZlLm9wZW5zaGlmdC5pby92MSIsCiAgImtpbmQiOiAiU3luY1NldCIsCiAgIm1ldGFkYXRhIjogewogICAgIm5hbWUiOiAic2FtcGxlIiwKICAgICJuYW1lc3BhY2UiOiAiYXJvLWY2MGFlOGEyLWJjYTEtNDk4Ny05MDU2LWYyZjZhMTgzN2NhYSIKICB9LAogICJzcGVjIjogewogICAgImNsdXN0ZXJEZXBsb3ltZW50UmVmcyI6IFtdLAogICAgInJlc291cmNlcyI6IFsKICAgICAgewogICAgICAgICJhcGlWZXJzaW9uIjogInYxIiwKICAgICAgICAia2luZCI6ICJDb25maWdNYXAiLAogICAgICAgICJtZXRhZGF0YSI6IHsKICAgICAgICAgICJuYW1lIjogIm15Y29uZmlnbWFwIgogICAgICAgIH0KICAgICAgfQogICAgXQogIH0KfQo=")
            .create();
    }
}
