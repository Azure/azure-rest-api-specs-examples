import com.azure.core.util.Context;

/** Samples for AutoProvisioningSettings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/GetAutoProvisioningSettingsSubscription_example.json
     */
    /**
     * Sample code: Get auto provisioning settings for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAutoProvisioningSettingsForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.autoProvisioningSettings().list(Context.NONE);
    }
}
