import com.azure.core.util.Context;

/** Samples for ConfigurationProfilesVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileVersion.json
     */
    /**
     * Sample code: Get a configuration profile version.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAConfigurationProfileVersion(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfilesVersions()
            .getWithResponse("customConfigurationProfile", "version1", "myResourceGroupName", Context.NONE);
    }
}
