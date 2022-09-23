import com.azure.resourcemanager.security.models.AutoProvision;

/** Samples for AutoProvisioningSettings Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/CreateAutoProvisioningSettingsSubscription_example.json
     */
    /**
     * Sample code: Create auto provisioning settings for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createAutoProvisioningSettingsForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.autoProvisioningSettings().define("default").withAutoProvision(AutoProvision.ON).create();
    }
}
