import com.azure.core.util.Context;

/** Samples for ConfigurationProfiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfile.json
     */
    /**
     * Sample code: Delete a configuration profile.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void deleteAConfigurationProfile(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfiles().deleteWithResponse("rg", "customConfigurationProfile", Context.NONE);
    }
}
