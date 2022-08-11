import com.azure.core.util.Context;

/** Samples for ConfigurationProfilesVersions ListChildResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileVersions.json
     */
    /**
     * Sample code: List configuration profile versions by configuration profile.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileVersionsByConfigurationProfile(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfilesVersions()
            .listChildResources("customConfigurationProfile", "myResourceGroupName", Context.NONE);
    }
}
