
/**
 * Samples for Secrets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/Secrets_CreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates a Secret with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void createsOrUpdatesASecretWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.secrets().define("childResourceName").withExistingOpenshiftcluster("resourceGroup", "resourceName")
            .create();
    }
}
