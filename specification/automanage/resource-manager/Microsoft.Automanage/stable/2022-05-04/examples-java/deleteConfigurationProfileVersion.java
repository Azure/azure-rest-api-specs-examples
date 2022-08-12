import com.azure.core.util.Context;

/** Samples for ConfigurationProfilesVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileVersion.json
     */
    /**
     * Sample code: Delete a configuration profile version.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void deleteAConfigurationProfileVersion(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfilesVersions()
            .deleteWithResponse("rg", "customConfigurationProfile", "version1", Context.NONE);
    }
}
