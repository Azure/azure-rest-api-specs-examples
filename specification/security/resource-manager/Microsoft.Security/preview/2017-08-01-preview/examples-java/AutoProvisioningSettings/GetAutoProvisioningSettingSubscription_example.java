/** Samples for AutoProvisioningSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/GetAutoProvisioningSettingSubscription_example.json
     */
    /**
     * Sample code: Get an auto provisioning setting for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getAnAutoProvisioningSettingForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.autoProvisioningSettings().getWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
