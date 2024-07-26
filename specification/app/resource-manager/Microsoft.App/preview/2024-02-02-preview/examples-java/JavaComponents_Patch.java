
import com.azure.resourcemanager.appcontainers.models.JavaComponent;
import com.azure.resourcemanager.appcontainers.models.JavaComponentConfigurationProperty;
import com.azure.resourcemanager.appcontainers.models.SpringBootAdminComponent;
import java.util.Arrays;

/**
 * Samples for JavaComponents Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/JavaComponents_Patch.json
     */
    /**
     * Sample code: Patch Java Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void patchJavaComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        JavaComponent resource = manager.javaComponents()
            .getWithResponse("examplerg", "myenvironment", "myjavacomponent", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(new SpringBootAdminComponent().withConfigurations(Arrays.asList(
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.ui.enable-toasts")
                    .withValue("true"),
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.monitor.status-interval")
                    .withValue("10000ms"))))
            .apply();
    }
}
