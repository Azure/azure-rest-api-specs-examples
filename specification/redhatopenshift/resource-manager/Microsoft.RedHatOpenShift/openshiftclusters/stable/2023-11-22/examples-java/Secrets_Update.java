
import com.azure.resourcemanager.redhatopenshift.models.Secret;

/**
 * Samples for Secrets Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/
     * examples/Secrets_Update.json
     */
    /**
     * Sample code: Updates a Secret with the specified subscription, resource group and resource name.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void updatesASecretWithTheSpecifiedSubscriptionResourceGroupAndResourceName(
        com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        Secret resource = manager.secrets()
            .getWithResponse("resourceGroup", "resourceName", "childResourceName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
