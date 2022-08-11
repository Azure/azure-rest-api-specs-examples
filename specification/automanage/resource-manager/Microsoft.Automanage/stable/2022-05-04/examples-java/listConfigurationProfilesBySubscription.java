import com.azure.core.util.Context;

/** Samples for ConfigurationProfiles List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfilesBySubscription.json
     */
    /**
     * Sample code: List configuration profiles by subscription.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfilesBySubscription(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfiles().list(Context.NONE);
    }
}
