
import com.azure.resourcemanager.appcontainers.models.DotNetComponentConfigurationProperty;
import com.azure.resourcemanager.appcontainers.models.DotNetComponentType;
import java.util.Arrays;

/**
 * Samples for DotNetComponents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DotNetComponents_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update .NET Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateNETComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().define("mydotnetcomponent")
            .withExistingManagedEnvironment("examplerg", "myenvironment")
            .withComponentType(DotNetComponentType.ASPIRE_DASHBOARD)
            .withConfigurations(Arrays.asList(
                new DotNetComponentConfigurationProperty().withPropertyName("dashboard-theme").withValue("dark")))
            .create();
    }
}
