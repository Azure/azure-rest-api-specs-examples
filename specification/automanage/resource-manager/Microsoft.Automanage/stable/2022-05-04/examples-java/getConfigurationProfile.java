import com.azure.core.util.Context;

/** Samples for ConfigurationProfiles GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfile.json
     */
    /**
     * Sample code: Get a configuration profile.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAConfigurationProfile(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfiles()
            .getByResourceGroupWithResponse("myResourceGroupName", "customConfigurationProfile", Context.NONE);
    }
}
