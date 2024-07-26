
import com.azure.resourcemanager.appcontainers.models.DotNetComponent;
import com.azure.resourcemanager.appcontainers.models.DotNetComponentConfigurationProperty;
import com.azure.resourcemanager.appcontainers.models.DotNetComponentType;
import java.util.Arrays;

/**
 * Samples for DotNetComponents Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/DotNetComponents_Patch.json
     */
    /**
     * Sample code: Patch .NET Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void patchNETComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        DotNetComponent resource = manager.dotNetComponents()
            .getWithResponse("examplerg", "myenvironment", "mydotnetcomponent", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withComponentType(DotNetComponentType.ASPIRE_DASHBOARD)
            .withConfigurations(Arrays.asList(
                new DotNetComponentConfigurationProperty().withPropertyName("dashboard-theme").withValue("dark")))
            .apply();
    }
}
